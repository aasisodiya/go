# Go Lang

- [Go Lang](#go-lang)
  - [What is Go?](#what-is-go)
  - [Variable Declaration in Go](#variable-declaration-in-go)
  - [Data Types in Go](#data-types-in-go)
  - [Sub String in String](#sub-string-in-string)
  - [Switch in Go](#switch-in-go)
  - [Important Points on Go](#important-points-on-go)

## What is Go?

- Go is an open-source programming language used to build simple, reliable and efficient software.
- Go was created by 3 People - Ken Thompson, Rob Pike and Robert Griesemer.
- Go is Simple
- Go is Strongly Typed (Variable has a type, and once declared it can't be assigned any other value type)
- Go is a procedural language with object oriented features.
- Go includes modern standard library
- Go Programs compiles to a single native binary
- Go Programs are fast because the code is native so you don't pay the performance penalty like with python, ruby or js.
- Go has garbage collection and is memory efficient. A well written go code can use upto 1/2 of that resource used by a Java Program.
- Go takes advantage of multiple cores. It has built-in concurrency.
- Well Known Projects That are written using Go - Hugo (Static Site Generator), Caddy (Web Server), Prometheus (Monitoring system & time series database), Docker, Kubernetes.

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

**Note**: You can not use short hand syntax out of function. For declaring package level variable var keyword is must. Type can be left out if it can be inferred. In Go all declared variable have a value. Also a declared variable needs to be used, else it will give you an error on compile.

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

## Sub String in String

You can very quickly get a substring of a string in golang using `[]`. `str[a,b]` where a is your starting index (included) and b is the ending index (excluded). Below code demonstrate the same.

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
