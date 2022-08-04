package main

import (
	"log"
	"regexp"
)

func main() {
	log.Println("Using regexp to remove numbers from string")
	sample := "ThisIsASampleStringWithNumbersFrom0To9InOrder0123456789AndThat'sIt"
	numbersRegExp := regexp.MustCompile("[0-9]+")
	op := numbersRegExp.ReplaceAllString(sample, "")
	log.Println(op)
}
