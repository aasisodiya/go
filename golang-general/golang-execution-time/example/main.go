package main

import (
	"fmt"
	"testing"

	"github.com/aasisodiya/go/benchmark/service"
)

func BenchmarkSampleFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		service.SampleFunction()
	}
}

func main() {
	fmt.Println("Ok")
	fmt.Println(testing.Benchmark(BenchmarkSampleFunction))
}
