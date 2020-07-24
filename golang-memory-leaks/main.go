package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"
)

var cnt int
var strarr2 []string

// Integrate below init() function properly in your code, plain copy-paste might cause issues so be careful
func init() {
	runtime.GOMAXPROCS(8)
	// Define a profilerURL for your program to use | you might sometime have to use different port if 8080 is not available
	profilerPort := ":8080"
	go func() {
		err := http.ListenAndServe(profilerPort, nil)
		if err != nil {
			fmt.Println("Error with ListenAndServe: ", profilerPort, err.Error())
			os.Exit(1)
		}
	}()
}

func main() {
	log.Println("Memory Leak Test Program")
	var strarr []string
	cnt = 0
	go Recur(strarr)
	// Just to allow program to keep running and call above go routine recur()
	time.Sleep(300 * time.Second) //Don't change this, It may crash your system
}

// Recur is a recursive function
func Recur(strarr []string) {
	cnt = cnt + 1
	time.Sleep(30 * time.Second) //Don't change this, It may crash your system
	strarr = append(strarr, time.Now().String())
	strarr2 = append(strarr2, time.Now().String())
	// This function's sole purpose is to act as a memory leak
	go Recur(strarr)
	Recur(strarr)
}
