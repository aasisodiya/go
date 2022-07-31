# Execute Command Using Go Lang

![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go.golang-execute-commands&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)

## Instructions

### Required Imports

```golang
import (
    "os/exec"
)
```

### Create Command You Want To Execute

For creating command we use `exec.Command()` with the arguments - command, followed by the elements of arg. Example given below. _Note: You can also configue buffered input and output here, please refer to this [link](https://golang.org/pkg/os/exec/)_

```go
// Command Name Only Example
// cmd := exec.Command("/home/ubuntu/programyouwanttorun")
// Command Name Only Example for Windows
cmd := exec.Command("c:/tmp/test.exe") //windows specific
// Command Name with Arguments Example
// args := ["args1","args2"]
// cmd := exec.Command("/home/ubuntu/programyouwanttorun", args...) // if arguments are needed
// Command Name with Arguments Example for Windows
cmd := exec.Command("c:/tmp/test.exe", "arg1", "arg2") //windows specific but with arguments
```

### Execute Above Built Command

Run/Execute the command with simple `cmd.Run()` and handle the error (if any)

```go
err := cmd.Run()
if err != nil {
    fmt.Printf("Error while executing:", err)
}
```

### Another Way To Run The Command

Here we are executing the command and also displaying its output. The Function `Output()` runs the command and returns its generated output. Any returned error will usually be of type \*ExitError. If c.Stderr was nil, Output populates ExitError.Stderr.

```go
out, err := exec.Command("c:/tmp/test.exe", "arg1", "arg2").Output()
if err != nil {
    fmt.Print(err)
}
fmt.Printf(string(out))
```

## Troubleshooting

- **Error:** `yourcommand: executable file not found in %path%`

  **Solution:** While I started working on this with Windows OS, I faced issue executing commands. So I found a workaround i.e to pass the complete path of the executable command. [Source](https://stackoverflow.com/questions/13008255/how-to-execute-a-simple-windows-command-in-golang)

  ```go
  c := exec.Command("cmd", "/C", "del", "D:\\a.txt")
  if err := c.Run(); err != nil {
      fmt.Println("Error: ", err)
  }
  ```

  ```go
  exec.Command("c:\\del.bat").Run()
  ```

## For Running the Sample Code

Please make sure to build the code in `test` folder first and then copy the same to `c:/tmp/` location for windows, else the referenced command in my code won't work as it is pointing to that program (OR) You can simply just update the code with your commands and execute it.

## Reference

- [Exec in Go Lang](https://golang.org/pkg/os/exec/)
- [How to execute a simple Windows command in Golang?](https://stackoverflow.com/questions/13008255/how-to-execute-a-simple-windows-command-in-golang)

[![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go&label=aasisodiya/go&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)](https://visitorbadge.io/status?path=aasisodiya.go)
