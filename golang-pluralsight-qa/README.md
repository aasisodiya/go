# Pluralsight Go Quiz

While Exploring Pluralsight I encountered some great set of Go Lang Questions in there quiz. Please do checkout [Pluralsight](https://www.pluralsight.com/) they have got some great features other than tutorials.

---

:octocat: How can a method be added to a built-in type, such as string?

:heavy_check_mark: Create a type alias for the built-in type and add the method to the alias type

:x: Methods cannot be added to built-in types

:x: No changes are required, the method can be defined as with any other type

:x: The method must be defined in a package called "builtin"

---

:octocat: Consider the following code, if changing line 3 to v := i.(int), what would be the output?

  1         func main() {
  2                 var i interface{} = "3"
  3                 v := i.(string)
  4                 fmt.Println(v)
  5        }
:x: nil

:heavy_check_mark: Panic error

:x: 3

:x: string

---

What is the difference between string literal and a raw string?

:x: You cannot create a multi line string with a string literal

:x: String literals occupy less stack space

:heavy_check_mark: Raw string preserves formatting and special characters

:x: Raw strings cannot be used in fmt.Sprintf

---

How do you define a struct which is deserialized from the following json?

{
  "myField": 1
}
:x: type MyStruct struct {
  myField int 
}
:x: type MyStruct struct {
  myField int `json`
}
:x: type MyStruct struct {
  Field int `json:myField`
}
:heavy_check_mark: type MyStruct struct {
  Field int `json:"myField"`
}

---

How would you compile a package and all of it's subpackages, placing the results in the workspace's pkg directory?

:x: Run

go install *

from within the parent package directory

:heavy_check_mark: Run

go install ./...

from within the parent package directory

:x: Run

go install ...

from within the parent package directory

:x: Run

go install **/*

from within the parent package directory

---

Consider that you have to write a function which converts any image format into JPEG. In order for an image to be converted from different formats, particular package(s) are required for its initialization (side effects), even if nothing is used from the package. Assume a PNG image format is being converted, which uses package "image/png". What is the correct way to import the package purely for its init function execution?

:heavy_check_mark: import _ “image/png”

:x: #import “image/png”

:x: import “image/png”

:x: import image/png

---

You have a project that has been using version 1.2 of a library at "https://github.com/foo/bar". Version 2 has just been released and you want to update to it, so you run the following command:
go get -u github.com/foo/bar
Unfortunately, your project does not use the new library when it is compiled. Why might this be?

:x: This command doesn't update existing packages

:heavy_check_mark: Another version is in the project's vendor directory

:x: The GOPATH hasn't been properly set

:x: The command requires the full URL, including protocol

---

Two projects, A and B, exist in your workspace that both depend on a third-party library. Due to the requirements of project A, the library is updated to the latest version. Unfortunately, the new version contains changes that break project B. What is the best way to address this, while having the least possible impact on the projects?

:x: Move the projects to separate workspaces and use separate GOPATHs for them

:heavy_check_mark: Move the library to project-specific vendor libraries

:x: Create a copy of the library with another name and use the copy for one of the projects

:x: Update project B to address the compatibility issue

---

Given a package square that has an init function, which of the following code samples of the init function is declared correct?


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
:x: 2

:heavy_check_mark: 1

:x: 4

:x: 3

---

What command can be used to display the help information about packages and commands?

:x: go man

:heavy_check_mark: go doc

:x: go info

:x: go help

---

In Golang 1.6 and above, local copies of external package dependencies are supported for vendoring. What is the correct way to vendor external package dependencies?

:x: Create a vendor folder in the $GOPATH/src folder and place all your dependencies in the vendor folder

:heavy_check_mark: Create a vendor folder in the worked on package’s root directory and place all package dependencies in the vendor folder

:x: Place all package dependencies in the worked on package’s root directory

:x: Place all package dependencies in $GOPATH/src folder

---

In Golang there is a convention when naming an interface with only one method. So consider the following interface:

type Generate interface {
    Generate() int
}
What would be an appropriate name to rename the Generate interface to comply with the convention?

:heavy_check_mark: Generator

:x: generate

:x: generator

:x: Gen

---

Which flag needs to be added in the command line to find a data race in the program and prints a report?

:x: -r

:x: -gorace

:x: -fsanitize

:heavy_check_mark: -race

---

 If you have to create a channel of strings, buffering up to 2 values. What is the correct way to declare it?

:x: messages := make(chan, string, 2)

:x: messages := make(string(chan, 2))

:heavy_check_mark: messages := make(chan string, 2)

:x: messages := make(chan(string, 2))

---

Given the following code snippet, in line 11 why does it not produce an output?



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
:x: Because sending a value into a messages channel on line 10 make the goroutine function terminate, and jumps to the next line after the goroutine function

:x: Because the goroutine function is missing a return statement at the end of the function

:heavy_check_mark: Because the messages channel is blocked at line 10 and won’t continue on until it receives a value from the channel

:x: Because sending a value into a messages channel on line 10 closes a messages channel and won’t continue on until reopens again

---

Consider a for loop which contains all three parts:

- an initialization variable
- a conditional expression
- a post statement

What must be true about the conditional expression in order for the enclosed code to be executed?

:x: The conditional expression must be evaluated to true, and the data type of the initialization variable must be integer, regardless of whether it is included within the conditional expression.

:x: The conditional expression must be evaluated to true, the data type of the initialization variable must be integer, and the initialization variable must be included within the expression.

:heavy_check_mark: The conditional expression must be evaluated to true, regardless of the data type of the initialization variable or whether it is included within the conditional expression.

:x: The conditional expression must be evaluated to true, regardless of the data type of the the initialization variable, but the initialization variable must be included within the expression.

---

Where should a call to `defer` be placed in order to ensure proper execution?

:x: It executes a function call before any of the surrounding function’s code is executed

:x: The defer statement must be in the middle to ensure proper execution when returning

:heavy_check_mark: It doesn't matter where a statement to execute defer is

:x: The defer statement must be at the top to ensure proper execution when returning

---

Consider that you are writing a small program and would like to log a file to detect races inside your current working directory $HOME/go/src/myprogram, rather than stderr. What would be the best way to make that work?

:x: Set the GOOS environment variable to "$HOME/go/src/myprogram"

:x: Set the GOTOOLDIR environment variable to "$HOME/go/src/myprogram"

:x: Set the GOPATH environment variable to "$HOME/go/src/myprogram"

:heavy_check_mark: Set the GORACE environment variable to "$HOME/go/src/myprogram"

---

