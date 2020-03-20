package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello from called program!")
	fmt.Println("Args that you passed are:", os.Args)
}