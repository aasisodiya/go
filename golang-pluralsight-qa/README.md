# Pluralsight Go Quiz

![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go.golang-pluralsight-qna&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)

While Exploring Pluralsight I encountered some great set of Go Lang Questions. Please do checkout [Pluralsight](https://www.pluralsight.com/) they have got some great features other than tutorials.

## Quiz

:grey_question: **How can a method be added to a built-in type, such as string?**

:heavy_check_mark: Create a type alias for the built-in type and add the method to the alias type

:heavy_multiplication_x: Methods cannot be added to built-in types

:heavy_multiplication_x: No changes are required, the method can be defined as with any other type

:heavy_multiplication_x: The method must be defined in a package called "builtin"

---

:grey_question: **Consider the following code, if changing line 3 to v := i.(int), what would be the output?**

```go
  1         func main() {
  2                 var i interface{} = "3"
  3                 v := i.(string)
  4                 fmt.Println(v)
  5        }
```

:heavy_multiplication_x: nil

:heavy_check_mark: Panic error

:heavy_multiplication_x: 3

:heavy_multiplication_x: string

---

:grey_question: **What is the difference between string literal and a raw string?**

:heavy_multiplication_x: You cannot create a multi line string with a string literal

:heavy_multiplication_x: String literals occupy less stack space

:heavy_check_mark: Raw string preserves formatting and special characters

:heavy_multiplication_x: Raw strings cannot be used in fmt.Sprintf

---

:grey_question: **How do you define a struct which is deserialized from the following json?**

```go
{
  "myField": 1
}
```

:heavy_multiplication_x:

```go
type MyStruct struct {
  myField int
}
```

:heavy_multiplication_x:

```go
type MyStruct struct {
  myField int `json`
}
```

:heavy_multiplication_x:

```go
type MyStruct struct {
  Field int `json:myField`
}
```

:heavy_check_mark:

```go
type MyStruct struct {
  Field int `json:"myField"`
}
```

---

:grey_question: **How would you compile a package and all of it's subpackages, placing the results in the workspace's pkg directory?**

:heavy_multiplication_x: Run

`go install *`

_from within the parent package directory_

:heavy_check_mark: Run

`go install ./...`

_from within the parent package directory_

:heavy_multiplication_x: Run

`go install ...`

_from within the parent package directory_

:heavy_multiplication_x: Run

`go install **/*`

_from within the parent package directory_

---

:grey_question: **Consider that you have to write a function which converts any image format into JPEG. In order for an image to be converted from different formats, particular package(s) are required for its initialization (side effects), even if nothing is used from the package. Assume a PNG image format is being converted, which uses package "image/png". What is the correct way to import the package purely for its init function execution?**

:heavy*check_mark: `import * “image/png”`

:heavy_multiplication_x: `#import “image/png”`

:heavy_multiplication_x: `import “image/png”`

:heavy_multiplication_x: `import image/png`

---

:grey_question: **You have a project that has been using version 1.2 of a library at "https://github.com/foo/bar". Version 2 has just been released and you want to update to it, so you run the following command:
`go get -u github.com/foo/bar`
Unfortunately, your project does not use the new library when it is compiled. Why might this be?**

:heavy_multiplication_x: This command doesn't update existing packages

:heavy_check_mark: Another version is in the project's vendor directory

:heavy_multiplication_x: The GOPATH hasn't been properly set

:heavy_multiplication_x: The command requires the full URL, including protocol

---

:grey_question: **Two projects, A and B, exist in your workspace that both depend on a third-party library. Due to the requirements of project A, the library is updated to the latest version. Unfortunately, the new version contains changes that break project B. What is the best way to address this, while having the least possible impact on the projects?**

:heavy_multiplication_x: Move the projects to separate workspaces and use separate GOPATHs for them

:heavy_check_mark: Move the library to project-specific vendor libraries

:heavy_multiplication_x: Create a copy of the library with another name and use the copy for one of the projects

:heavy_multiplication_x: Update project B to address the compatibility issue

---

:grey_question: **Given a package square that has an init function, which of the following code samples of the init function is declared correct?**

```go
package square

var side int


func init() {                      (1)
    side = 4
}


func init(x int) {                 (2)
    side = x
}


func init() int {                  (3)
    return 4
}


func init(x int) int {             (4)
    return x
}
```

:heavy_multiplication_x: 2

:heavy_check_mark: 1

:heavy_multiplication_x: 4

:heavy_multiplication_x: 3

---

:grey_question: **What command can be used to display the help information about packages and commands?**

:heavy_multiplication_x: `go man`

:heavy_check_mark: `go doc`

:heavy_multiplication_x: `go info`

:heavy_multiplication_x: `go help`

---

:grey_question: **In Golang 1.6 and above, local copies of external package dependencies are supported for vendoring. What is the correct way to vendor external package dependencies?**

:heavy_multiplication_x: Create a vendor folder in the $GOPATH/src folder and place all your dependencies in the vendor folder

:heavy_check_mark: Create a vendor folder in the worked on package’s root directory and place all package dependencies in the vendor folder

:heavy_multiplication_x: Place all package dependencies in the worked on package’s root directory

:heavy_multiplication_x: Place all package dependencies in $GOPATH/src folder

---

:grey_question: **In Golang there is a convention when naming an interface with only one method. So consider the following interface:**

```go
type Generate interface {
    Generate() int
}
```

**What would be an appropriate name to rename the Generate interface to comply with the convention?**

:heavy_check_mark: Generator

:heavy_multiplication_x: generate

:heavy_multiplication_x: generator

:heavy_multiplication_x: Gen

---

:grey_question: **Which flag needs to be added in the command line to find a data race in the program and prints a report?**

:heavy_multiplication_x: `-r`

:heavy_multiplication_x: `-gorace`

:heavy_multiplication_x: `-fsanitize`

:heavy_check_mark: `-race`

---

:grey_question: **If you have to create a channel of strings, buffering up to 2 values. What is the correct way to declare it?**

:heavy_multiplication_x: `messages := make(chan, string, 2)`

:heavy_multiplication_x: `messages := make(string(chan, 2))`

:heavy_check_mark: `messages := make(chan string, 2)`

:heavy_multiplication_x: `messages := make(chan(string, 2))`

---

:grey_question: **Given the following code snippet, in line 11 why does it not produce an output?**

```go
  1         package main
  2
  3         import "fmt"
  4
  5         func main() {
  6
  7                         messages := make(chan int)
  8
  9                         go func() {
 10                                 messages <- 0
 11                                 fmt.Print("Hello")
 12                         }()
 13
 14                         fmt.Println("World")
 15         }
```

:heavy_multiplication_x: Because sending a value into a messages channel on line 10 make the goroutine function terminate, and jumps to the next line after the goroutine function

:heavy_multiplication_x: Because the goroutine function is missing a return statement at the end of the function

:heavy_check_mark: Because the messages channel is blocked at line 10 and won’t continue on until it receives a value from the channel

:heavy_multiplication_x: Because sending a value into a messages channel on line 10 closes a messages channel and won’t continue on until reopens again

---

:grey_question: **Consider a for loop which contains all three parts:**

- an initialization variable
- a conditional expression
- a post statement

**What must be true about the conditional expression in order for the enclosed code to be executed?**

:heavy_multiplication_x: The conditional expression must be evaluated to true, and the data type of the initialization variable must be integer, regardless of whether it is included within the conditional expression.

:heavy_multiplication_x: The conditional expression must be evaluated to true, the data type of the initialization variable must be integer, and the initialization variable must be included within the expression.

:heavy_check_mark: The conditional expression must be evaluated to true, regardless of the data type of the initialization variable or whether it is included within the conditional expression.

:heavy_multiplication_x: The conditional expression must be evaluated to true, regardless of the data type of the the initialization variable, but the initialization variable must be included within the expression.

---

:grey_question: **Where should a call to `defer` be placed in order to ensure proper execution?**

:heavy_multiplication_x: It executes a function call before any of the surrounding function’s code is executed

:heavy_multiplication_x: The defer statement must be in the middle to ensure proper execution when returning

:heavy_check_mark: It doesn't matter where a statement to execute defer is

:heavy_multiplication_x: The defer statement must be at the top to ensure proper execution when returning

---

:grey_question: **Consider that you are writing a small program and would like to log a file to detect races inside your current working directory $HOME/go/src/myprogram, rather than stderr. What would be the best way to make that work?**

:heavy_multiplication_x: Set the GOOS environment variable to "$HOME/go/src/myprogram"

:heavy_multiplication_x: Set the GOTOOLDIR environment variable to "$HOME/go/src/myprogram"

:heavy_multiplication_x: Set the GOPATH environment variable to "$HOME/go/src/myprogram"

:heavy_check_mark: Set the GORACE environment variable to "$HOME/go/src/myprogram"

## Reference

[Pluralsight](https://app.pluralsight.com/score/skill-assessment/go/intro?context=skills#/v2/landing)

[![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go&label=aasisodiya/go&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)](https://visitorbadge.io/status?path=aasisodiya.go)
