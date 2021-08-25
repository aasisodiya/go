package main

import (
	"fmt"

	"github.com/aasisodiya/go/codecoverage/service"
)

func main() {
	// This is just a sample code to check/demonstrate code coverage
	fmt.Println("This is a simple example of Code Coverage in GoLang")
	cube := service.NewCube(10)
	fmt.Println(cube.SurfaceArea())
	fmt.Println(cube.Volume())
}
