# Lambda using Go Lang

> **Important:** Below document is specific to Windows Developer

## Steps to deploy Go Program on Lambda

1. Build the code (Ex. `$Env:GOOS = "linux"; go build -o main`)
1. Zip the executable (Ex. `~\Go\Bin\build-lambda-zip.exe --output main.zip main`)
1. Upload to Lambda
1. Change Handler to `main` (if not done already)
1. Save the lambda changes
1. Create a sample event to test the code
1. Test run the code and check for output

All steps details are given below

## For developers on Windows

Windows developers may have trouble producing a zip file that marks the binary as executable on Linux. To create a .zip that will work on AWS Lambda, the `build-lambda-zip` tool may be helpful.

1. Get the tool

    ``` shell
    set GO111MODULE=on
    go.exe get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip
    ```

1. Use the tool from your `GOPATH`. If you have a default installation of Go, the tool will be in `%USERPROFILE%\Go\bin`. 

    * in cmd.exe:

        ``` bat
        set GOOS=linux
        set GOARCH=amd64
        set CGO_ENABLED=0
        go build -o main main.go
        %USERPROFILE%\Go\bin\build-lambda-zip.exe --output main.zip main
        ```

    * in Powershell:

        ``` posh
        $env:GOOS = "linux"
        $env:GOARCH = "amd64"
        $env:CGO_ENABLED = "0"
        go build -o main main.go
        ~\Go\Bin\build-lambda-zip.exe --output main.zip main
        ```

### To do a cross-compilation in single line

> Prefer using Powershell for windows (as it is default in VS Code)

```powershell
$Env:GOOS = "linux"; go build -o main

OR

$Env:GOOS = "linux"; $Env:GOARCH = "amd64"; go build
```

To specify the output file name:

```powershell
$Env:GOOS = "linux"; $Env:GOARCH = "amd64"; go build -o outputfilename
```

### Sample Lambda Event

```json
{
  "Resource": "test",
  "Path": "test",
  "HTTPMethod": "GET",
  "Body": "test"
}
```

Use above sample if you are testing your lambda with API Gateway as trigger. Also be cautious while creating test event. It is strictly related to request object. Above test event is created by referring to struct of [events.APIGatewayProxyRequest.](https://github.com/aws/aws-lambda-go/blob/master/events/apigw.go)

### Sample Lambda Response

```json
{
  "statusCode": 200,
  "headers": null,
  "multiValueHeaders": null,
  "body": "Ok here is GET Method"
}
```

Even Response is strictly related to your response object. Above response is in reference with [events.APIGatewayProxyResponse.](https://github.com/aws/aws-lambda-go/blob/master/events/apigw.go)

## Troubleshooting

* **Issue:** Lambda gives following error while testing

    ```json
    {
      "errorMessage": "fork/exec /var/task/main: permission denied",
      "errorType": "PathError"
    }
    ```

    **Solution:** Issue commonly occurs for windows users. You can't zip the executable file directly instead are required to use command depending on what you are using:

  * Command Prompt

    ```bat
    %USERPROFILE%\Go\bin\build-lambda-zip.exe --output main.zip main
    ```

  * Powershell:

    ```posh
    ~\Go\Bin\build-lambda-zip.exe --output main.zip main
    ```

* **Issue:** Facing issue while moving executable file to zip: `flag provided but not defined: -o`

    ```bat
    Reference\GoLang\Lambda-GoLang>%USERPROFILE%\Go\bin\build-lambda-zip.exe -o main.zip main
    Incorrect Usage. flag provided but not defined: -o

    NAME:
       build-lambda-zip - Put an executable and supplemental files into a zip file that works with AWS Lambda.

    USAGE:
       build-lambda-zip.exe [global options] command [command options] [arguments...]

    COMMANDS:
       help, h  Shows a list of commands or help for one command

    GLOBAL OPTIONS:
       --output value  output file path for the zip. Defaults to the first input file name.
       --help, -h      show help (default: false)
    flag provided but not defined: -o
    ```

  * **Solution:** Use `--output` instead of `-o`

      So your command becomes like this for CMD: `%USERPROFILE%\Go\bin\build-lambda-zip.exe --output main.zip main` and for Powershell: `~\Go\Bin\build-lambda-zip.exe --output main.zip main`

## Instructions to Write Lambda in Golang (Essentials)

1. Define Imports

    ```golang
    import (
      "errors"

      "github.com/aws/aws-lambda-go/events"
      "github.com/aws/aws-lambda-go/lambda"
    )
    ```

    * Above 3 imports are necessary for most cases, `"github.com/aws/aws-lambda-go/events"` is required for creating Request and Response object for lambda
    * `"errors"` is required for handeling errors for invalid request

1. Create a function to handle request with events.APIGatewayProxyRequest as parameter and events.APIGatewayProxyResponse and error as return values. Then you can handle the request based on HTTPMethod type (a check for type is good idea). Also make sure to handle the invalid requests. For building response we use [`events.APIGatewayProxyResponse`](https://github.com/aws/aws-lambda-go/blob/master/events/apigw.go) and then return the same.

    ```golang
    func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
      if request.HTTPMethod == "GET" {
        stringResponse := "Ok here is GET Method"
        APIResponse := events.APIGatewayProxyResponse{Body: stringResponse, StatusCode: 200}
        return APIResponse, nil
      } else {
        err := errors.New("Method Not Allowed")
        APIResponse := events.APIGatewayProxyResponse{Body: "Method Not Allowed!", StatusCode: 502}
        return APIResponse, err
      }
    }
    ```
1. Now call above function from `main()` function using `lambda.Start()` with `HandleRequest` as input parameter.

    ```golang
    func main() {
      lambda.Start(HandleRequest)
      // lambda.Start(Handler)
    }
    ```

## Reference

* [Issue: /var/task/main: permission denied](https://github.com/awslabs/aws-sam-cli/issues/274)
* [AWS Lambda in Go Lang for Windows Developer](https://github.com/aws/aws-lambda-go#for-developers-on-windows)
* [Events Structure in Go Lang for Lambda](https://godoc.org/github.com/aws/aws-lambda-go/events)
* [How to Cross-Compile Go Program on Windows](https://stackoverflow.com/questions/50911153/how-to-crosscompile-go-programs-on-windows-10)
