package main

import (
	"fmt"
	"pkg1"
	"pkg2"
)

func main() {
	fmt.Println("test")
	pkg1.Pkg1Test()
	pkg2.Pkg2Test()
}

func test1(){
	fmt.Println("main")
}

func Test1(){
	fmt.Println("main")
}
