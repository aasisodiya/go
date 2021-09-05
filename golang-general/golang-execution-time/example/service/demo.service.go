package service

import (
	"fmt"
)

// SampleFunction is just for sample
func SampleFunction() (cnt int) {
	fmt.Println("Sample Function is Called")
	cnt = 0
	for i := 0; i < 10; i++ {
		cnt++
	}
	return cnt
}
