# Execution Time in Go Lang

![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go.golang-general.execution-time&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)

- [Execution Time in Go Lang](#execution-time-in-go-lang)
  - [Simple Code Logic for Execution Time](#simple-code-logic-for-execution-time)
  - [Measuring Function Execution Time](#measuring-function-execution-time)
    - [Shared Code Logic](#shared-code-logic)
    - [Implementation Code Logic](#implementation-code-logic)
  - [Benchmarking in Go Lang](#benchmarking-in-go-lang)

In order to measure performance or benchmarking a go program you can measure it's execution time.

## Simple Code Logic for Execution Time

Below is a simple go code that measures its execution time.

```go
import (
    "log"
    "time"
)

func main() {
    start := time.Now()
    // Your Code Logic....
    log.Println(time.Since(start))
}
```

In above code what we have done is recorded a start time. Then after the code logic is executed we are printing the elapsed time w.r.t the recorded start time.

## Measuring Function Execution Time

In order to measure the execution time for a function in go lang, the easier way is to use `defer` and some shared code logic.

### Shared Code Logic

```go
// Track function return the message along with start time for the function/method
func Track(msg string) (string, time.Time) {
    return msg, time.Now()
}

// Duration function prints out the message along with time elapsed since start
func Duration(msg string, start time.Time) {
    log.Printf("%v: %v\n", msg, time.Since(start))
}
```

### Implementation Code Logic

Now you can use the above shared code logic inside your function. Below is a sample code for reference. We use defer here since we want execution time to be calculated and displayed on end of the function/method execution

```go
// SomeFunction is just a dummy sample function
func SomeFunction() {
    defer Duration(Track("SomeFunction"))
    time.Sleep(1 * time.Minute)
}
```

## Benchmarking in Go Lang

Benchmarks are written in `*_test.go` file and the Benchmarking Methods start with `Benchmark*`.

> For demo purpose you can use [this code](https://github.com/aasisodiya/go/tree/master/golang-general/golang-execution-time/example)

Sample code looks like given below

```go
package service

import (
    "testing"
)

func BenchmarkSampleFunction(b *testing.B) {
    for i := 0; i < b.N; i++ {
        SampleFunction()
    }
}
```

Instead of using `go test`, you can also call `testing.Benchmark` from a command

```go
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
    fmt.Println(testing.Benchmark(BenchmarkSampleFunction))
}
```
