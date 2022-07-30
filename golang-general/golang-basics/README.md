# Go Lang

![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go.golang-general.basics&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)

- [Go Lang](#go-lang)
  - [What is Go?](#what-is-go)
  - [Why Go?](#why-go)
  - [Go Commands](#go-commands)
  - [Variable Declaration in Go](#variable-declaration-in-go)
  - [Data Types in Go](#data-types-in-go)
  - [Naming Convention in Go](#naming-convention-in-go)
  - [Sub String in String](#sub-string-in-string)
  - [Switch in Go](#switch-in-go)
  - [Important Points on Go](#important-points-on-go)
  - [Random Number in Go](#random-number-in-go)
  - [Arrays in Go](#arrays-in-go)
  - [Slices in Go](#slices-in-go)
  - [Map in Go](#map-in-go)
  - [Struct in Go](#struct-in-go)
  - [Struct Tags](#struct-tags)
  - [Methods in Go](#methods-in-go)
  - [Define Methods on Type](#define-methods-on-type)
  - [Interfaces in Go](#interfaces-in-go)
    - [Better example of Interfaces in Go](#better-example-of-interfaces-in-go)
  - [Type Assertion Vs Conversion](#type-assertion-vs-conversion)
  - [Type Switch](#type-switch)
  - [Errors in Go](#errors-in-go)
  - [CSP & Go Routines](#csp--go-routines)
  - [Channels for Sending Data between Go Routines](#channels-for-sending-data-between-go-routines)
  - [Buffer Channel](#buffer-channel)
  - [Deadlock in Go Routine](#deadlock-in-go-routine)
  - [Closing a channel](#closing-a-channel)
  - [Using Channel with for loop](#using-channel-with-for-loop)
  - [Select in Go](#select-in-go)
  - [Testing in Go](#testing-in-go)
  - [Printing out an Object](#printing-out-an-object)
  - [Pointer in Go](#pointer-in-go)
  - [Reference](#reference)

![Visitor](https://visitor-badge.glitch.me/badge?page_id=aasisodiya.go)

## What is Go?

- Go is an open-source programming language used to build simple, reliable and efficient software.
- Go was created by 3 People - Ken Thompson, Rob Pike and Robert Griesemer.
- Go is Simple
- Go is Compiled Language, its not interpreted language
- Go is Strongly Typed (Variable has a type, and once declared it can't be assigned any other value type)
- Go is a procedural language with object oriented features.
- Go includes modern standard library
- Go Programs compiles to a single native binary
- Go Programs are fast because the code is native so you don't pay the performance penalty like with python, ruby or js.
- Go has garbage collection and is memory efficient. A well written go code can use upto 1/2 of that resource used by a Java Program.
- Go takes advantage of multiple cores. It has built-in concurrency.
- Well Known Projects That are written using Go - Hugo (Static Site Generator), Caddy (Web Server), Prometheus (Monitoring system & time series database), Docker, Kubernetes.

---

## Why Go?

- Go is compiled language hence its much faster compared to interpreted language
- No Runtime dependency, hence version dependency is eliminated
- Go compiles to a single binary file
- Strongly Typed
- Built-in Concurrency
- Cross Platform
- Easier than any other language

---

> Click on this [link to download Go](https://go.dev/dl/)

---

## Go Commands

| Command       | Description                                                                              |
| ------------- | ---------------------------------------------------------------------------------------- |
| `go build`    | Will compile the code and create an executable file                                      |
| `go run`      | Will compile and run the code                                                            |
| `go run .`    | Will compile and run the code for multiple\* go files on windows                         |
| `go run *.go` | Will compile and run the code for multiple\* go files on macos                           |
| `go fmt`      | Will format the code                                                                     |
| `go install`  | Will compile and install the package                                                     |
| `go get`      | command to import any external public package. Example `go get github.com/aasisodiya/go` |
| `go test`     | Will execute the tests                                                                   |
| `go mod init` | Will initialize a module. Example `go mod init github.com/aasisodiya/goprgm`             |

`* here multiple go files refer to multiple package main files in same folder, ex main.go, handler.go are in same folder and have same package main`

> Only main package on build produces an executable file, and building a package other than main gives no executable file. Basically package main is executable package.

> List of all awesome packages in go [Click Here](https://github.com/avelino/awesome-go)

## Variable Declaration in Go

Below are the different ways to declare a variable in go

```go
// using type and int
var i int = 10
// var inferred
var i = 10
// using colon inferred
i := 10
```

**Note**: You can not use short hand syntax out of function. Type can be left out if it can be inferred. In Go all declared variable have a value. Also a declared variable _(inside a function)_ needs to be used, else it will give you an error on compile. For declaring package level variable `var` keyword is must. For Example, below code will give you an error : `Non-declaration statement outside function body`

```go
package main

import "fmt"

pi := 3.14 // wrong
// var pi = 3.14 or var pi float64 = 3.14

func main() {
 fmt.Println(pi)
}
```

---

## Data Types in Go

Following are list of types in go

- bool
- string
- int int8 int16 int32 int64
- uint uint8 uint16 uint32 uint64 uintptr
- byte // alias for uint8
- rune // alias for int32 and represents a Unicode code point
- float32 float64
- complex64 complex128

---

## Naming Convention in Go

- Anything (Function/Method/Variable) written or defined with camelCase is set to private access
- Anything (Function/Method/Variable) written or defined with TitleCase is set to public access

---

## Sub String in String

You can very quickly get a substring of a string in golang using `[]`. `str[a,b]` where a is your starting index (included) and b is the ending index (excluded). Below code demonstrate the same. And for a valid case `a <= b` i.e a needs to be less than equal to b.

```go
package main
import (
  "fmt"
)
func main() {
  str := "Everything is Awesome"
  str1 := str[0:10] // will print characters from 0 to 9 (and not 10) i.e. Everything
  str2 := str[:10] // will also print characters from start to 9 (and not 10) i.e. Everything
  str3 := str[14:21] // will print characters from 14 to 20 (and not 21) i.e. Awesome
  // str3 := str[14:len(str)] // will also print characters from 14 to end i.e. Awesome
  str4 := str[14:] // will also print characters from 14 to end i.e. Awesome
  str5 := str[0:21] // will print the whole string i.e. Everything is Awesome
  str6 := str[0:len(str)] // will also print the whole string i.e. Everything is Awesome
  str7 := str[:] // will also print the whole string i.e. Everything is Awesome
  str8 := str[11:13] // will print the characters from index 11 to 12 (not 13) i.e. is
  fmt.Printf(`str = %s, str1 = %s, str2 = %s, str3 = %s, str4 = %s,
  str5 = %s, str6 = %s, str7 = %s, str8 = %s, len(str) = %d`,
  str, str1, str2, str3, str4, str5, str6, str7, str8, len(str))
}
```

> [Click here](https://play.golang.org/p/bOcEQwgSnAM) to open/edit/run above code

---

## Switch in Go

```go
switch word {
  case "hey":
    // do something
  default:
    // do something
}
```

In go switch you don't have to use break, unlike other programming languages where case statements are executed until break. But in go every case has its own block and doesn't need a break to be specified. But if you want it to behave like other programming languages then you can use a keyword `fallthrough` which will help you do the same.

```go
switch word {
  case "hey":
    // do something then fallthrough
    fallthrough
  case "hello":
    // do something
  case "hi":
    // do something
  default:
    // do something
}
```

If you want to do something same for 2 different cases then just separate them using comma as given in below example

```go
switch word {
  case "hey", "hi":
    // do something
  default:
    // do something
}
```

---

## Important Points on Go

- Go doesn't support ([Optional or Named Parameters](https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/classes-and-structs/named-and-optional-arguments))
- Go doesn't support function overloading (i.e. you cannot have 2 function with same name declared within same scope, even if they have different number/types of parameters)
- Use `_` i.e. underscore to ignore a value
- All function calls in go are call by value i.e. a copy of value of variable is passed to the function, and not the reference to variable itself. Example to demo the same is given below

  ```go
  package main

  import (
    "fmt"
  )

  func test(a int, str string, arr [2]int) {
    a = a * 2
    for i := 0; i < len(arr); i++ {
      arr[i] *= 2
    }
    str += "test"
    fmt.Println("inside function:", a, str, arr)
  }
  func main() {
    a := 89
    str := "this is "
    arr := [2]int{2, 3}
    test(a, str, arr)
    fmt.Println("inside main:", a, str, arr)
  }
  ```

  Above program output is given below, [Click Here](https://gt/p/NyC3s0oE2Mr) to open/run/edit the code

  ```txt
  inside function: 178 this is test [4 6]
  inside main: 89 this is  [2 3]
  ```

- You can assign a function to a variable. For example in above code you can do the following and it will still work.

  ```go
  newfunc := test
  // and now use newfunc instead of test
  newfunc(a, str, arr)
  ```

  [Click Here](https://play.golang.org/p/JDic0kCJZqD) to open/run/edit the code

- You cannot declare a function inside another function, but you can always use anonymous function inside a function. Example of the same is given below

  ```go
  package main

  import (
    "fmt"
  )

  func main() {
    //func test() string {
    //return "Test"
    //}
    // Above code will give you compile error
    anonymous := func() string {
      return "Anonymous"
    }
    fmt.Println(anonymous())
  }
  ```

  [Click Here](https://play.golang.org/p/I35K7Jcf3_E) to open/run/edit the code

- You can pass function as argument as well, for ex. `func test(f func(int)int)` here function `test` will take any function with given signature i.e. that take input `int` and return `int` ex. `func foo(int a) int`. Another best example is `http.HandleFunc("/welcome", printWelcome)`
- Relative path package import is not considered as good practice.

---

## Random Number in Go

For getting/generating a random number in go you can use `"math/rand"` package. Example for the same is given below

```go
package main

import (
  "fmt"
  "math/rand"
)

func main() {
  rand.Seed(1)
  fmt.Println("Random Number", rand.Intn(10)) // 1
  fmt.Println("Random Number", rand.Intn(10)) // 7
}
```

> [Click Here](https://play.golang.org/p/xfBMX9_3kON) to edit/run the code

---

## Arrays in Go

Declaring an Array

```go
var myArray [10]int
```

Initializing an Array

```go
var arr1 = [5]int{} // will initialize an array with 0 i.e [0,0,0,0,0]
var arr2 = [5]int{1,2,3} // will initialize an array of size 5 filled with values given in {} and then followed by 0 i.e [1,2,3,0,0]
var arr2 = [5]int{1,2,3,4,5} // will initialize an array of size 5 i.e [1,2,3,4,5]
var arr3 = […]int{1,2,3,4}   // the compiler will count the array elements for you
```

---

## Slices in Go

Slice is a growable sequence of values of a single specified type. You can create a slice with slice literal i.e. same as an array but without specifying size, or by using `make([]T, len, cap) []T` function where T is type, len is length, cap is capacity.

Declaring a Slice

```go
var intSlice []int
```

Initialize a Slice

```go
var intSlice = []int{1,2,3} // [1 2 3]
len(intSlice) // 3
cap(intSLice) // 3
// OR
var intSlice = make([]int, 2, 4) // [0, 0, nil, nil] i.e. that will just print [0 0]
len(intSlice) // 2
cap(intSLice) // 4
```

Example of Slice is given below

```go
package main

import (
  "fmt"
)

func main() {
  exampleSlice := make([]string, 0) // this will create an empty slice
  // exampleSlice := []string{}
  exampleSlice = append(exampleSlice, "Akash") // so when you append a string size become 1
  fmt.Println(len(exampleSlice)) // 1
  exampleSlice = append(exampleSlice, "Akash")
  fmt.Println(len(exampleSlice)) // 2
  exampleSlice2 := make([]string, 4) // but this will create a slice with size 4 with empty strings
  exampleSlice2 = append(exampleSlice2, "Akash") // so when you append a string it get added to
  // existing 4 elements and thus size become 5
  fmt.Println(len(exampleSlice2)) // 5
}
```

> [Click Here](https://play.golang.org/p/IrGzSo4fiuj) to run/edit the code

---

## Map in Go

A map is a data type that associates a value of one data type to a value of another data type. You can create a map using `make`. You should avoid declareing map as `var someMap map[string]int`. Map definition is `map[indexType]valueType`.

```go
m := make(map[string]int)
// OR
m := map[string]int{}
m := map[string]int{
  "a":1,
  "b":2,
  "c":3,
}
```

- Order of iteration in map is completely random, you can not rely on map for its value to be in a same order.
- To add a value to map you can simply do `m["d"]=4` and to Remove a value from map using `delete` i.e `delete(m, "b")`

| `sampleMap := map[int]int{}`          | `sampleMap := make(map[int]int)`                                     |
| ------------------------------------- | -------------------------------------------------------------------- |
| Can be used to create a non empty map | Creates an empty map                                                 |
| We cannot specify the capacity here   | We can specify capacity, which can help to reduce future allocations |

> Above data is taken from [link](https://stackoverflow.com/questions/16959992/creating-map-with-without-make)

---

## Struct in Go

First of all Structs are not Objects!, as there is no inheritance in structs. They might be similar. To create struct we use keyword `struct` as given in below example

```go
package main

import (
  "fmt"
)

type Sample struct {
  PublicProperty  string
  privateProperty int
}

type StructWithFunc struct {
  SomeFunction    func(int) string
}

func main() {
  s0 := Sample{}            // Init with default values
  s1 := Sample{"Value1", 8} // For such initialization all fields must be specified
  s2 := Sample{             // For such initialization, you don't need to specify every field
    privateProperty: 90,
  }
  s2.PublicProperty = "Value"
  s2.privateProperty = 900
  fmt.Println(s0, s1, s2)
  // { 0} {Value1 8} {Value 900}
  sf := StructWithFunc{test}
  sf.SomeFunction(2)
}

func test(i int) string {
  fmt.Println("Test")
  return "ok"
}

```

> [Click Here](https://go.dev/play/p/wYC-3_7mU0u) to run above code

> You can also have function inside a struct

Example of embedded struct is given below

```go
package main

import (
  "fmt"
)

type SimpleStruct struct {
  A string
  B int
}

type AnotherStruct struct {
  C string
  D SimpleStruct
}

type EmbeddedStruct struct {
  A string
  SimpleStruct
}

func main() {
  s := SimpleStruct{A: "a string", B: 8}
  as := AnotherStruct{C: "c string", D: s}
  es := EmbeddedStruct{A: "em string", SimpleStruct: s}
  fmt.Println(s, as, es)
  fmt.Println(as.D.A)
  fmt.Println(es.A)
  fmt.Println(es.B) // Using embedded struct you can directly access the property of embedded struct
  // But what about property A?? as it is similar to that of property in EmbeddedStruct
  // For such scenario you can use the the struct name, example give below
  fmt.Println(es.SimpleStruct.A)
  fmt.Println(es.SimpleStruct.B)
}
```

> [Click Here](https://play.golang.org/p/9p2pcGbp5VF) to edit/run above code

---

## Struct Tags

Go allows to put metadata on fields of struct, which is used for marshalling and unmarshalling data from into/out of struct. `json:"name"` is called a struct tag

```go
package main

import (
  "encoding/json"
  "fmt"
)

type Car struct {
  Name string `json:"name"`
  Cost int    `json:"cost"`
}

func main() {
  car1 := `{
    "name":"Audi R8",
    "cost":150000
  }`
  var audi Car
  json.Unmarshal([]byte(car1), &audi)
  fmt.Println(audi)
  audiCar, err := json.Marshal(audi)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(string(audiCar))
}
```

> [Click Here](https://play.golang.org/p/V8ZH798HQJ2) to edit/run the code

---

## Methods in Go

Method declaration has a method receiver declared between keyword `func` and the name of the method. Method receiver looks like parameter declaration. For example `func (s *Sample) SomeMethod()` in here `(s *Sample)` is the receiver. Most receiver in the go code are of pointer types, and it is referred as best practices in go documentation.

```go
package main

import (
  "fmt"
)

type Car struct {
  Name string `json:"name"`
  Cost int    `json:"cost"`
}

func (c *Car) DisplayData() {
  fmt.Println(fmt.Sprintf("%s has cost of %d dollars", c.Name, c.Cost))
}

func (c *Car) ChangeName(newName string) {
  c.Name = newName
  fmt.Println("Name Updated to ", newName)
}

func (c Car) ChangeCost(newCost int) {
  c.Cost = newCost
  fmt.Println("Cost Updated to ", newCost)
}

func main() {
  audi := Car{"Audi R8", 150000}
  audi.DisplayData() // Audi R8 has cost of 150000 dollars
  audi.ChangeName("Audi R8 V10") // Name Updated to  Audi R8 V10
  audi.DisplayData() // Audi R8 V10 has cost of 150000 dollars
  audi.ChangeCost(200000) // Cost Updated to  200000
  audi.DisplayData() // Audi R8 V10 has cost of 150000 dollars
}
```

> [Click Here](https://play.golang.org/p/AsN8-3Dn0Mh) to edit/run above code

You might notice that name changes but cost doesn't that is because, for method `ChangeName` we used reference receiver i.e. `(c *car)` whereas for `ChangeCost` method we use value receiver i.e. `(c Car)`. So when you use `(c Car)` you're actually creating a copy and updating its value. But when you use `(c *car)` you are actually referencing the actual property and thus it gets updated correctly.

**In Conclusion use value receiver when you are not modifying the value of the struct and use reference receiver when you are modifying the value of the struct**.

> By Convention, usually the receiver parameter name is single or double letter as given in above example for Car we have used only `c` hence `c *Car`. `this` or `self` aren't defined/used in go. Go basically avoids any mentions of `this` or `self`, and using the same can be breaking convention.

---

## Define Methods on Type

In go you can define methods on type. Example of the same is given below

```go
package main

import (
  "fmt"
)

type myInt int

func (mi myInt) isEven() bool {
  return mi%2 == 0
}

func main() {
  m := myInt(10)
  fmt.Println(m)
  fmt.Println(m.isEven())
}
```

> [Click Here](https://play.golang.org/p/zH44BsRKwYz) to edit/run above code

---

## Interfaces in Go

A Type that lists a set of methods but provides no implementation. Sample code is given below

```go
package main

import (
  "fmt"
)

type Sample interface {
  Test(int) bool
}

type A struct {
  num1 int
  num2 int
}

type B struct {
  num1 int
}

func (b B) Test(i int) bool {
  return b.num1 > i
}

func (a A) Test(i int) bool {
  return a.num1 > i && a.num2 < i
}

func main() {
  samples := []Sample{
    A{10, 2},
    B{7},
  }
  for _, sample := range samples {
    fmt.Println(sample.Test(5))
    fmt.Println(sample.Test(8))
  }
}
```

> [Click here](https://play.golang.org/p/6i6O4krk2Q3) to edit/run the code

### Better example of Interfaces in Go

Interfaces in go helps you reuse the code that is common for different type of variables.

Now lets take an example to demonstrate better use of interface. In below example you can see that Print Area method defined on both Rectangle and Square struct are the similar logic. Hence the interface solved our issue by allowing us to reuse the code. `PrintArea(a Area)` is an replacement for both of them.

```go
package main

import "fmt"

func main() {
  // Declaring the AreaCalculator to validate that the code fulfills the interface contract.
  var r, s AreaCalculator
  r = Rectangle{
    Width:  10,
    Height: 20,
  }
  s = Square{
    Side: 20,
  }
  PrintArea(r)
  PrintArea(s)
}

func (r Rectangle) GetArea() int {
  return r.Height * r.Width
}
func (s Square) GetArea() int {
  return (s.Side * s.Side)
}

// Below methods are not needed
// func (r *Rectangle) PrintArea() {
//   fmt.Println(r.Area)
// }
// func (s Square) PrintArea() {
//   fmt.Println(s.Area)
// }

type Rectangle struct {
  Height int
  Width  int
  Area   int
}

type Square struct {
  Side int
  Area int
}

type AreaCalculator interface {
  GetArea() int
}

func PrintArea(a AreaCalculator) {
  fmt.Println(a.GetArea())
}
```

> [Click here](https://go.dev/play/p/vtSzNQDsWuW) to edit/run the code

- Interface helps us to make sure that a particular type qualifies for an interface when it implements all the methods mentioned in that interface.
- We cannot create a values out of interface. As its not a concrete type like int, string, etc.
- Interfaces are not generic type
- Interfaces are implicit i.e we don't have to link any type with interfaces explicitly
- Interfaces are a contract to help us manage types
- In order for a struct to be used with the interface, it must validate the contract i.e the struct should implement all the methods mentioned in the interface.

---

## Type Assertion Vs Conversion

In Type Conversion, you are changing the type i.e. int to float and so on. But for Type Assertion you are revealing the underlying type and stripping away the interface that wraps it. You can only do type assertion on an interface. You can do type conversion on any type.

Empty Interface in Go, i.e. with no methods at all. Any types could match this.

```go
package main

import (
  "fmt"
)

func main() {
  var i interface{}
  i = "string"
  j := i.(string)
  fmt.Println(i, j)
  k, ok := i.(int)
  fmt.Println(k, ok)
  // m := i.(int) this will panic as its wrong assertion
}
```

> [Click here](https://play.golang.org/p/6jhVv1XJgo6) to edit/run the code

When doing type assertion you need to be sure that type is correct else code will panic.

---

## Type Switch

`i.(type)` such code outside of switch block will given an error. Sample code of type switch is given below.

```go
package main

import (
  "fmt"
)

func typeCheck(i interface{}) {
  switch i := i.(type) {
  case int:
    fmt.Println("Its Integer:", i)
  case string:
    fmt.Println("Its String:", i)
  default:
    fmt.Println("Type Unknown:", i)
  }
}

func main() {
  var i interface{}
  i = "string"
  typeCheck(i) // Its String: string
  i = 1000
  typeCheck(i) // Its Integer: 1000
  i = 10.0
  typeCheck(i) // Type Unknown: 10
}
```

> [Click Here](https://play.golang.org/p/4HRcJG4fAOw) to edit/run the code

---

> `type testFunc func(int) bool` is also a valid code ([ref](https://stackoverflow.com/questions/9398739/working-with-function-types-in-go)) this simply says that testFunc represents type `func(int) bool`. Let say for example you use `type myInt int` will create a new type `myInt` which directly represents `int`. click here for better example - [Define Methods on Type](#define-methods-on-type)
> Interface is only nil if there is no underlying value assigned to it

---

## Errors in Go

- There are no exceptions in Go
- By convention the last return value is usually error
- You just have to put in extra code to make sure there are all required validations and error checks and nothing gets skipped
- New Error can be defined using `errors.New()`
- Most usefull errors package are `"errors"` and `"github.com/pkg/errors"`

---

## CSP & Go Routines

- Usually concurrency in most languages are defined using threads as unit of execution. Each thread is mapped to os-level thread. They share data data by acquiring locks on shared pieces of data. Many languages have concurrency support in their libraries.
- Concurrency in Go is different and its concurrency model is based on **CSP (Communicating Sequential Process)** and is built into the language.
- Go Routines are lightweight processes managed by the Go Runtime
- Go Routines are not threads, threads are managed by OS while the go routines are managed by go runtime scheduler.
- Go Runtime Schedules go routines across threads automatically
- The Go Runtime increases the number of threads at start-up and schedule go routines across this threads automatically.
- Go Routines have several benefits:
  - Go Routine creation is faster than thread creation (Why? because you are not creating os-level resource)
  - Go Routine stack sizes are smaller than thread stack sizes and can grow as needed
  - Switching between go routines is faster than switching between threads.
  - Due to all this features go can spawn 1000 and 10000 of go routines (in a single process)

Sample Go Routine Code is given below

```go
package main

import (
  "fmt"
)

func test() {
  fmt.Println("Go Routine executed me")
}

func main() {
  go test()
}
```

> [Click here](https://play.golang.org/p/tB51QPoIsk5) to edit/run the code

When you run the above code you will notice that you don't get any output. That is because main go routine didn't knew it had to wait for the other goroutine to finish execution. So for our go routine to execute we have different ways:

- By Pausing using `time.Sleep(1*time.Second)`, **This is not recommended in real programs and should be avoided**
- By using Wait Group, for which we use `"sync"` package
- Try to pass shared variables in argument to the go routine function in order to avoid issues related to referencing.

```go
package main

import (
  "fmt"
  "sync"
)

func test() {
  fmt.Println("Go Routine executed me")
}

func main() {
  var wg sync.WaitGroup
  // We don't have to initialize it, just declare it
  wg.Add(1)
  // This tells the wait group that we have added 1 go routine to run
  // Now we launch a closure as a go routine
  go func() {
    test()
    // Now we call Done method to notify that go routine has completed
    wg.Done()
  }()
  // Now here at last we call Wait on wait group
  wg.Wait()
  // Wait helps us to make sure that go routine finish execution
}
```

> [Click Here](https://play.golang.org/p/AAJy3FTQbl6) to edit/run above code

For best practices try keeping concurrency logic different from business logic.

Now lets have a look at go routine with data, First code uses direct reference to the data. Thus You won't get the output you expect, because all go routine share the same variable. If you run `go vet` for below code you will get following error: `./prog.go:18:9: loop variable i captured by func literal`

```go
package main

import (
  "fmt"
  "sync"
)

func test(i int) {
  fmt.Println("Go Routine executed me: ", i)
}

func main() {
  var wg sync.WaitGroup
  // direct reference (You won't get the output you expect, because all go routine share the same variable)
  for i := 0; i < 100; i++ {
    wg.Add(1)
    go func() {
      test(i)
      wg.Done()
    }()
  }
  wg.Wait()
}
```

> [Click Here](https://play.golang.org/p/r3jNmOb22IY) to edit/run above code

If you want to pass different parameter to each go routine, You should pass i as parameter to the go routine. As shown in below code

```go
package main

import (
  "fmt"
  "sync"
)

func test(i int) {
  fmt.Println("Go Routine executed me: ", i)
}

func main() {
  var wg sync.WaitGroup
  // direct reference (You won't get the output you expect, because all go routine share the same variable)
  for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(x int) {
      test(x)
      wg.Done()
    }(i)
  }
  wg.Wait()
}
```

> [Click Here](https://play.golang.org/p/uKosu1pilSb) to edit/run above code

---

## Channels for Sending Data between Go Routines

- Channel is another built-in type in Go which is used for transferring data between go routines
- Multiple Go Routines can write to a single channel and Multiple Go Routines can read from a single channel
- Data on channel is typed. ex. `make(chan string)` is used to create a channel of typed string
- By Default channel reads and writes are synchronous. i.e. when a go routines writes to a channel it pauses until another go routines read from that channel and vice versa i.e. a go routine pauses when it read from channel and resume only when there is data available to read.
- Always close / defer close your channels, by using `defer close(channelName)`

Sample code for channel is given below

```go
package main

import (
  "fmt"
)

func main() {
  in := make(chan string)
  defer close(in)
  out := make(chan string)
  defer close(out)

  go func() {
    name := <-in                         // Reading a value from channel
    out <- fmt.Sprintf("Hello, " + name) // Writing a value to a channel
  }()
  in <- "Akash"    // Writing a value to a channel
  message := <-out // Reading a value from a channel
  fmt.Println(message)
}
```

> [Click Here](https://play.golang.org/p/etpjJucNXoV) to edit/run above code

## Buffer Channel

- If you do not want to wait for a channel to be read after you write to it. For such scenario we can use buffer channel.
- A buffer channel lets you write a specified number of times before there is a read on the channel
- A buffer is never infinite
- If you fill out a buffer, your write will pause the go routine until there is a read on the channel from another go routine.
- Example: `pool := make(chan func(int)int, workers)`

```go
package main

import (
  "fmt"
)

func main() {
  out := make(chan int, 10)
  defer close(out)
  for i := 0; i < 10; i++ {
    go func(localI int) {
      out <- localI
    }(i)
  }
  var result []int
  for i := 0; i < 10; i++ {
    val := <-out
    result = append(result, val)
  }
  fmt.Println(result)
}
```

> [Click Here](https://play.golang.org/p/x5cd4d3YuYE) to edit/run above code

---

## Deadlock in Go Routine

Sample code for dead lock in go routine is given below

```go
package main

import (
  "fmt"
)

func main() {
  in := make(chan int)
  defer close(in)
  out := make(chan int)
  defer close(out)
  go func() {
    for {
      i := <-in
      out <- i * 2
    }
  }()
  in <- 1
  in <- 2
  o1 := <-out
  o2 := <-out
  fmt.Println(o1, o2)
}
```

> [Click Here](https://play.golang.org/p/9bjuxXxS2VS) to edit/run above code

When you run above code you will get following output, showing a deadlock has occurred

```text
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
  /tmp/sandbox2823182645/prog.go:19 +0x136

goroutine 6 [chan send]:
main.main.func1()
  /tmp/sandbox2823182645/prog.go:15 +0x55
created by main.main
  /tmp/sandbox2823182645/prog.go:12 +0x10f

Program exited: status 2.
```

Now to solve above issue you have an option of making channel in/out as buffered channel of size greater than or equal to 2. [Click here](https://play.golang.org/p/yuf6EmIE_zr) to run this scenario. But this solution is very specific to above scenario. So another better solution will be to make sure that you read the channel before you write any other value to the channel. [Click here](https://play.golang.org/p/nrZtSBATwhH) to run this scenario

---

## Closing a channel

When there is no more data to write to a channel, then only the channel can be closed. Some important points on closing a channel:

- Closing a channel does not vipes out its content
- Closing the channel means that it will not have any more values written to it
- Zero value for a channel is nil
- If a buffered channel is closed, any value in buffer is still available for it to be read.
- You must write to a closed channel / call close second time on a closed channel. If you do any then your program will panic.
- If you read from a closed channel with no more data then read returns with zero value of the type of the channel. Ex. for channels of int we get 0, for string we get empty string.

But then how to know if the value is from a closed channel? for that we using `val, ok := <-chan` idiom. If value ok is true then channel is still open and if the value is false then channel is closed.

| Operation | Unbuffered                       | Buffered                  | nil          | closed                                                              |
| --------- | -------------------------------- | ------------------------- | ------------ | ------------------------------------------------------------------- |
| read      | Pause until something is written | pause if buffer is empty  | Hang forever | Return immediately w/zero value (use comma-ok to see if its closed) |
| write     | Pause until something is read    | Pause only if buffer full | Hang forever | PANIC                                                               |
| close     | Works                            | Works                     | PANIC        | PANIC                                                               |

---

## Using Channel with for loop

```go
for _, v := range myChannel{
  // Logic
}
```

Above code is similar to below code

```go
for {
  rec := <-myChannel
  // Logic
}
```

## Select in Go

- Control Structure for Concurrency in Go
- You also reach a deadlock if you write to a channel but there is no read on it or vice versa.
- select is the key to ensure non blocking reads and writes from channel. Simply put your channel reads and write with a default.
- We use done channel pattern with select

---

## Testing in Go

- Testing in go is not like mocha, selenium, etc.
- For creating a test you need to create a new file ending with `_test.go`
- `go test` command will execute all the tests in given package
- Basic testing function is like `func TestSomeFunction(t *testing.T)`

---

## Printing out an Object

Below is the simple way of printing an object with all it properties and their respective values.

```go
fmt.Println("%+v", objectName)
```

---

## Pointer in Go

In below code, `*SomeObject` is a type description, which in below case is an pointer to `SomeObject`. Whereas `(*pointerToObject)` here `*` is an operator, representing value of the pointer its representing.

```go
func (pointerToObject *SomeObject) update(){
  (*pointerToObject).Property = 9
}
```

Above and below function does the same thing

```go
func (pointerToObject *SomeObject) update(){
  pointerToObject.Property = 9
}
```

| Note             |            |
| ---------------- | ---------- |
| Address to Value | `*address` |
| Value to Address | `&value`   |

Sample program for explaining pointers in go

```go
package main

import "fmt"

type SomeType struct {
  Property int
}

func main() {
  t := SomeType{Property: 42}
  fmt.Printf("%+v\n", t)
  // You can directly call the method on the struct even though it is not an pointer as given in method definition below
  t.UpdateProperty(10)
  fmt.Printf("%+v\n", t)
  tptr := &t
  // You can also call the method on the pointer to the struct as given in method definition below
  tptr.UpdateProperty2(20)
  fmt.Printf("%+v\n", t)
  // As you can see both way works in go language, i.e you can call a method on the struct or pointer to the struct.
  // Same holds true for the property update as given in method UpdateProperty2 where we are using (*t).
  t.UpdateProperty2(30)
  fmt.Printf("%+v\n", t)
}

func (t *SomeType) UpdateProperty(a int) {
  t.Property = a
}

func (t *SomeType) UpdateProperty2(a int) {
  (*t).Property = a
}
```

When passing slice in the function there is a direct reference to that slice and not like in other cases where copies are created and updated. For example

```go
func UpdateSlice(slc []string) {
  slc[0] = "Hola!"
}

mySlice := []string{"Hey","There"}
UpdateSlice(mySlice)
// this above function call will update the slice from [Hey There] to [Hola There]
```

Hence when you pass int, string, float or struct in any function in go, it creates a copy of that argument and then use them inside the function.

> fmt.Println(\*&val) is similar to fmt.Println(val)

| Value Types | Reference Types |
| ----------- | --------------- |
| int         | slices          |
| float       | maps            |
| string      | channels        |
| bool        | pointers        |
| structs     | functions       |

For above `Refernce Types` you don't have to worry about pointers. But for `Value Types` you have to use pointers to change them in functions.

## Reference

- [Effective Go](https://golang.org/doc/effective_go)
- [Standard Tools](https://golang.org/doc/cmd)
- [Learn Go in 3 Hours](https://www.udemy.com/course/learn-go-in-3-hours/)

[![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)](https://visitorbadge.io/status?path=aasisodiya.go)
