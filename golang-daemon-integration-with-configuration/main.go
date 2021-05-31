// Template / Example Code
package main

import (
	"fmt"
	"github.com/takama/daemon"
	"gopkg.in/yaml.v2"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	description = "some description of the program"
	// This needs to be configured properly with a static file address (don't use relative address here)
	configFile = "/home/ubuntu/configs/config.yml"
)

//    dependencies that are NOT required by the service, but might be used
var dependencies = []string{}

var stdlog, errlog *log.Logger

// Service has embedded daemon
type Service struct {
	daemon.Daemon
}

var cfg Config

// Config struct for config file
type Config struct {
	Service struct {
		Name  string        `yaml:"name"`
		Port  string        `yaml:"port"`
		Timer time.Duration `yaml:"timer"`
	} `yaml:"service"`
	LogFile struct {
		Location string `yaml:"location"`
		Name     string `yaml:"name"`
	} `yaml:"logfile"`
}

// Manage by daemon commands or run the daemon
func (service *Service) Manage() (string, error) {
	log.Println("Inside Manager Function")
	usage := "Usage: myservice install | remove | start | stop | status"

	// if received any kind of command, do it
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return service.Install()
		case "remove":
			return service.Remove()
		case "start":
			return service.Start()
		case "stop":
			return service.Stop()
		case "status":
			return service.Status()
		default:
			return usage, nil
		}
	}

	// Do something, call your goroutines, etc

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Set up listener for defined host and port
	tcpport := ":" + cfg.Service.Port
	log.Println("TCP Port : ", tcpport)
	listener, err := net.Listen("tcp", tcpport)
	if err != nil {
		return "Possibly was a problem with the port binding", err
	}

	// set up channel on which to send accepted connections
	listen := make(chan net.Conn, 100)
	go acceptConnection(listener, listen)
	log.Println("Time Period : ", cfg.Service.Timer*time.Millisecond*1000)
	ticker := time.NewTicker(cfg.Service.Timer * time.Millisecond * 1000)
	tickerDone := make(chan bool)
	// loop work cycle with accept connections or interrupt
	// by system signal
	for {
		select {
		case <-ticker.C:
			stdlog.Println("yourMethod called")
			go yourMethod() //call your target method here, which you want to trigger
		case <-tickerDone:
			return "Timer stopped", nil
		case killSignal := <-interrupt:
			stdlog.Println("Got signal:", killSignal)
			stdlog.Println("Stoping listening on ", listener.Addr())
			listener.Close()
			if killSignal == os.Interrupt {
				return "Daemon was interruped by system signal", nil
			}
			return "Daemon was killed", nil
		}
	}

	// never happen, but need to complete code
	return usage, nil
}

func yourMethod() {
	fmt.Println("Ticked", time.Now())
	log.Println("Ticked", time.Now())
}

// Accept a client connection and collect it in a channel
func acceptConnection(listener net.Listener, listen chan<- net.Conn) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		listen <- conn
	}
}

func handleClient(client net.Conn) {
	for {
		buf := make([]byte, 4096)
		numbytes, err := client.Read(buf)
		if numbytes == 0 || err != nil {
			return
		}
		client.Write(buf[:numbytes])
	}
}

func init() {
	stdlog = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	errlog = log.New(os.Stderr, "", log.Ldate|log.Ltime)
}

func processFatalError(err error) {
	fmt.Println(err)
	fmt.Println(errlog) // Log to system Log  /var/log/syslog
	os.Exit(2)
}

func readFile(cfg *Config) {
	f, err := os.Open(configFile)
	if err != nil {
		processFatalError(err)
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processFatalError(err)
	}
}

func main() {
	readFile(&cfg)
	logFileName := cfg.LogFile.Location + string(os.PathSeparator) + cfg.LogFile.Name
	log.Println(cfg.Service.Name)
	log.Println(cfg.Service.Port)
	log.Println(cfg.Service.Timer)
	fmt.Println(logFileName)
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer logFile.Close()
	log.SetOutput(logFile)
	// New ----------------
	log.Println("Main()")
	srv, err := daemon.New(cfg.Service.Name, description, daemon.SystemDaemon, dependencies...)
	if err != nil {
		errlog.Println("Error: ", err)
		os.Exit(1)
	}
	service := &Service{srv}
	status, err := service.Manage()
	if err != nil {
		errlog.Println(status, "\nError: ", err)
		os.Exit(1)
	}
}
