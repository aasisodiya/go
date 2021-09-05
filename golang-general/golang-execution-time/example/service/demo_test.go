package service

import (
	"testing"
)

func BenchmarkSampleFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SampleFunction()
	}
}
