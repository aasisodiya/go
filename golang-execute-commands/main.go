package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Command with output
	out, err := exec.Command("c:/tmp/test.exe", "arg1", "arg2").Output()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf(string(out))

	// Build a command and then execute but it won't show output
	// cmd := exec.Command("/home/ubuntu/go/programs/programyouwanttorun")
	// cmd := exec.Command("c:/tmp/test.exe") //windows specific
	cmd := exec.Command("c:/tmp/test.exe", "arg1", "arg2") //windows specific but with arguments
	// args := ["args1","args2"]
	// cmd := exec.Command("/home/ubuntu/programyouwanttorun", args...) // if arguments are needed
	fmt.Println(cmd)
	err1 := cmd.Run()
	if err1 != nil {
		fmt.Printf("Error while executing:", err1)
	}
}
