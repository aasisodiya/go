# Lambda using Go Lang

![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go.golang-aws-lambda&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)

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

## Q&A

### How to convert your Go struct to JSON

**Solution:** Use `json.Marshal()` with import `"encoding/json"`

```go
package main

import (
    "encoding/json"
    "fmt"
)

// Temp is a sample
type Temp struct {
    ID   int
    Name string `json:"myName,omitempty"`
}

func main() {
    temp := &Temp{
        ID:   1,
        Name: "Akash", //this comma is required
    }
    data, err := json.Marshal(temp)
    fmt.Println(string(data), err)
}
// Output:
// {"ID":1,"myName":"Akash"} <nil>
```

`json.Marshal()` method takes `temp` i.e our struct as input and generates output into `data` if everything goes well else return an error in `err`. Also data that we get in `data` is of type []byte. So for printing the value we use `string()` method.

Also you might have noticed `json:"myName,omitempty"` being used in above code. Now what this does is, while generating JSON it replaces the property name with json defined name. That's why our generated output looks like this `{"ID":1,"myName":"akash"}` and not like this `{"ID":1,"Name":"akash"}`. Also `omitempty` helps us if we want to elimate empty data from our generated json. So if the value for temp.Name was undefined then our result will be like this `{"ID":1}`

**Important Point to Note:** If you don't put a comma just before closing bracket `}` of object definition you will get an error. *Please refer to comment in code above*

>**Reference:** <br> [Stackoverflow](https://stackoverflow.com/questions/8270816/converting-go-struct-to-json) <br> [Go Lang Official Doc](https://golang.org/pkg/encoding/json/#Marshal)

---

### How to hide fields while using `json.Marshal()` to create JSON Object

**Solution:** Use `json:"-"`

Consider you have a struct with 8 properties and when you build its JSON response you will get all those 8 properties in your JSON response object. But if you want to remove some properties without much hustle you can simply use `json:"-"` after you define the property type.

```go
type Temp struct {
    ID       string `json:"id"`
    Name     string
    NickName string
}
```

If you use above given struct, then on using `json.Marshal()` you will get below result

```json
{
    "id":       "aasisodiya", // ID got replaced with id because of `json:"id"`
    "Name":     "Akash",      // Name and NickName are as it is as we didn't
    "NickName": "Bhai"        // specify any formatting for them in json
}
```

Now lets eliminate NickName from our JSON result for that we will use `json:"-"`

```go
type Temp struct {
    ID       string `json:"id"`
    Name     string
    NickName string `json:"-"`
}
```

Now what `json:"-"` does is, it removes that property while marshalling. Hence we get following Result

```json
{
    "id":       "aasisodiya",
    "Name":     "Akash"
}
```

> **Reference:** <br> [Stackoverflow](https://stackoverflow.com/questions/17306358/removing-fields-from-struct-or-hiding-them-in-json-response)

---

### How to build a test event for your Go Lang Lambda

For building a sample test event for your lambda, you will have to refer to below struct of APIGatewayProxyRequest. Since it is the object that carries request data to our go program.

```go
type APIGatewayProxyRequest struct {
    Resource                        string                        `json:"resource"` // The resource path defined in API Gateway
    Path                            string                        `json:"path"`     // The url path for the caller
    HTTPMethod                      string                        `json:"httpMethod"`
    Headers                         map[string]string             `json:"headers"`
    MultiValueHeaders               map[string][]string           `json:"multiValueHeaders"`
    QueryStringParameters           map[string]string             `json:"queryStringParameters"`
    MultiValueQueryStringParameters map[string][]string           `json:"multiValueQueryStringParameters"`
    PathParameters                  map[string]string             `json:"pathParameters"`
    StageVariables                  map[string]string             `json:"stageVariables"`
    RequestContext                  APIGatewayProxyRequestContext `json:"requestContext"`
    Body                            string                        `json:"body"`
    IsBase64Encoded                 bool                          `json:"isBase64Encoded,omitempty"`
}
```

Now while refering to above struct you will make sure to list out all the properties that you have used in your code. For example consider you have created a `request` object of type `APIGatewayProxyRequest` and used `request.PathParameters` and `request.Body` somewhere in your code then **it's vital that your test event must include this properties else your test will fail.**

So your test event will look like this

```json
{
    "pathParameters": {
        "pathParameter1": "pathParameterValue1",
        "pathParameter2": "pathParameterValue2"
    },
    "body": "sample text body"
}
```

Now you can see that we have used respective properties' equivalent json mapping for creating a test event. For example `PathParameters` is defined using its json representation from `APIGatewayProxyRequest` struct i.e `pathParameters`. Similarly for `Body` we used `body`. Like this you will do the same for all your required properties in test event.

---

### Adding element to an array of variable size

**Solution:** Use `func append(s []T, vs ...T) []T`

Consider you have to populate some data into an array but you are uncertain of the array size. Then you can simply start by defining the array and then using append to add elements to an array. Refer the example below

```go
// Temp is a sample
type Temp struct {
    ID   int
    Name string
}

func main(){
    // Lets define an array of Temp
    var tempArray []Temp
    // now use append to add an element
    tempArray = append(tempArray, Temp{5,"Akash"})
    // You can more than 1 elements at once too
    tempArray = append(tempArray, Temp{6,"Aakash"}, Temp{7,"Aditya"})
}
```

While using `append()` you pass the Array as first parameter and second parameter onwards are your element you want to add. You can add multiple elements as well.

> **Reference:** <br> [Tour Go Lang](https://tour.golang.org/moretypes/15) <br> [Stackoverflow](https://stackoverflow.com/questions/3387273/how-to-implement-resizable-arrays-in-go/3437599)

Note: *If You want to run your sample go code online you can use below mentioned link*

> **Go Online IDE:** [The Go Playground](https://play.golang.org/)

[![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)](https://visitorbadge.io/status?path=aasisodiya.go)
