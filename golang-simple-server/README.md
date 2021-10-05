# Simple Go Lang Server

- [Simple Go Lang Server](#simple-go-lang-server)
  - [Simple Web Server in Go](#simple-web-server-in-go)
  - [Powershell Command to Test](#powershell-command-to-test)

## Simple Web Server in Go

You can create a simple web server in go using `"net/http"` package. In which you can use `http.HandleFunc()` to handle your request.

```go
package main

import(
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
    name := req.URL.Query().Get("name")
    message := fmt.Sprintf("Hello %s", name)
    byteResponse := []byte(message)
    w.Write(byteResponse)
}

func main() {
    pattern := "/hello"
    http.HandleFunc(pattern, hello)
    http.ListenAndServe(":8080", nil)
}
```

Command to Test

```bash
# Powershell Command
$response = Invoke-RestMethod 'http://localhost:8080/hello?name=Akash' -Method 'GET' -Headers $headers
$response | ConvertTo-Json

# Command Prompt Command
curl http://localhost:8080/hello?name=Akash
```

## Powershell Command to Test

Below powershell command will help you test the code written in [main.go](https://github.com/aasisodiya/go/blob/master/golang-simple-server/main.go)

```ps
$response = Invoke-RestMethod 'http://localhost:8090/get-sample' -Method 'GET' -Headers $headers
$response | ConvertTo-Json
```
