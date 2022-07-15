# AWS DynamoDB Operations in Go Lang

![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go.golang-aws-operations.aws-dynamodb&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)

## Steps Followed While Creating a Shared DynamoDB Operation Function

* Create Session

    ```golang
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-west-2")},
    )
    ```

* Create DynamoDB client

    ```golang
    svc := dynamodb.New(sess)
    ```

OR

* Create DynamoDB Client with Session

    ```golang
    svc := dynamodb.New(session.New(&aws.Config{Region: aws.String("us-west-2")},))
    ```

* Call required DynamoDB Operation

    ```golang
    result, err := svc.GetItem(&dynamodb.GetItemInput{
        TableName: aws.String("ENTER_YOUR_TABLE_NAME_HERE"),
        Key: map[string]*dynamodb.AttributeValue{
            "NUM_KEY": {
                N: aws.String("NUM_VALUE"),
            },
            "STRING_KEY": {
                S: aws.String("STRING_VALUE"),
            },
        },
    })
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    ```

    or

    ```golang
    input := &dynamodb.GetItemInput{
        TableName: aws.String("ENTER_YOUR_TABLE_NAME_HERE"),
        Key: map[string]*dynamodb.AttributeValue{
            "NUM_KEY": {
                N: aws.String("NUM_VALUE"),
            },
            "STRING_KEY": {
                S: aws.String("STRING_VALUE"),
            },
        },
    }
    result, err := svc.GetItem(input)
    if err != nil {
        if aerr, ok := err.(awserr.Error); ok {
            switch aerr.Code() {
            case dynamodb.ErrCodeProvisionedThroughputExceededException:
                fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
            case dynamodb.ErrCodeResourceNotFoundException:
                fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
            case dynamodb.ErrCodeRequestLimitExceeded:
                fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
            case dynamodb.ErrCodeInternalServerError:
                fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
            default:
                fmt.Println(aerr.Error())
            }
        } else {
            // Print the error, cast err to awserr.Error to get the Code and
            // Message from an error.
            fmt.Println(err.Error())
        }
        return
    }
    fmt.Println(result)
    ```

## Steps followed for creating Lambda

* Create Go Module

    ```powershell
    go mod init github.com/aasisodiya/dynammodb
    ```

* Add below line to go.mod, where github.com/aasisodiya/aws is replaced with relative path of aws folder containing the required functions

    ```mod
    replace github.com/aasisodiya/aws => ../../aws
    ```

* Create main.go and import shared dynamodb package and all necessary imports (example given below)

    ```golang
    import (
        "fmt"

        "encoding/json"

        dynamodbop "github.com/aasisodiya/aws/dynamodb"
        "github.com/aws/aws-lambda-go/events"

        "github.com/aws/aws-lambda-go/lambda"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/service/dynamodb"
    )
    ```

* Create Main Function and call `lambda.Start(HandleRequest)`

    ```golang
    func main() {
        lambda.Start(HandleRequest)
    }
    ```

* Create HandleRequest Function, which will contain 3 important parts

    1. Extract data from Request Body (Make sure to return proper error if required request data is missing)
    2. Use the request data to build a input parameter for your DynamoDB Operation
    3. Perform required DynamoDB Operation
    4. Check Result and Errors properly and return the same, you might have to convert the Result according to events.APIGatewayProxyResponse object

    Below is the example for GetItem Operation on DynamoDB

    ```golang
    // HandleRequest Method (Using Headers)
    func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
        if request.HTTPMethod == "GET" {
            tableName := request.Headers["tableName"]
            region := request.Headers["region"]
            primarykey := request.Headers["primarykey"]
            // Check if all required values are received in header
            if tableName == "" || region == "" || primarykey == "" {
                APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"Headers Parameters: [tableName, region, primarykey] all values are required\"}", StatusCode: 400}
                return APIResponse, nil
            }

            // Build the Query Attribute
            key := map[string]*dynamodb.AttributeValue{
                "pkey": {
                    S: aws.String(primarykey),
                },
            }

            // Query the table for record
            fmt.Println("Getting record from table", tableName, " in Region", region, " for key:", primarykey)
            result, err := dynamodbop.GetItem(region, tableName, key)

            // Check for errors
            if err != nil {
                fmt.Println(err.Error())
                APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"" + err.Error() + "\"}", StatusCode: 502}
                return APIResponse, nil
            }

            // Covert the result to JSON
            fmt.Print(result)
            json, err := json.Marshal(result)

            // Return the results (convert json to string)
            APIResponse := events.APIGatewayProxyResponse{Body: string(json), StatusCode: 200, IsBase64Encoded: false}
            return APIResponse, nil
        }

        // Return error if Method Type is different
        APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"Method Not Allowed!\"}", StatusCode: 405}
        return APIResponse, nil
    }
    ```

    **Something that I noticed:** the json for result houses a very large data which if returned directly in JSON format to response will show up like this to API caller

    ```json
    {
        "ConsumedCapacity": null,
        "Item": {
            "key1": {
                "B": null,
                "BOOL": null,
                "BS": null,
                "L": null,
                "M": null,
                "N": null,
                "NS": null,
                "NULL": null,
                "S": "value1",
                "SS": null
            },
            "pkey": {
                "B": null,
                "BOOL": null,
                "BS": null,
                "L": null,
                "M": null,
                "N": null,
                "NS": null,
                "NULL": null,
                "S": "pkey1",
                "SS": null
            }
        }
    }
    ```

    It looks bulky and has quite many null values, that's why we can use `dynamodbattribute.UnmarshalMap(result.Item, &item)` method, where `item` is the variable created from a struct TestItem. What this method does is accepts a map[string]*dynamodb.AttributeValue and converts it to a interface{} or struct.

    First you have to create a struct for your returning Dynamodb Item, (**Note:** In case the struct differs from returning item , then that value from db will be dropped so be cautious)

    ```go
    // TestItem !!! this struct won't work (Even though AWS docs use this)
    type TestItem struct {
        key1 string
        pkey string
    }
    // TestItem is a struct for your dynamodb item (This works!)
    type TestItem struct {
        Key1 string `json:"key1"`
        Pkey string `json:"pkey"`
    }
    ```

    Then you have to use `dynamodbattribute.UnmarshalMap` method

    ```go

    // Unmarshal the result to a struct
    item := TestItem{}

    err = dynamodbattribute.UnmarshalMap(result.Item, &item)
    if err != nil {
        panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
    }

    // Covert the result to JSON
    json, err := json.Marshal(item)
    ```

    Now this json value is what you can use in response, which will now show up as below to API caller

    ```json
    {
    "key1": "value1",
    "pkey": "pkey1"
    }
    ```

* Finally compile and zip the code, then upload it to Lambda

    ```powershell
    $Env:GOOS = "linux"; go build -o main
    ~\Go\Bin\build-lambda-zip.exe --output main.zip main
    ```

    > If you face any issue in above commands then You can refer to documentation of golang-aws-lambda/README.md for more details

## Troubleshooting

* `Error: cannot convert result (type *"github.com/aws/aws-sdk-go/service/dynamodb".GetItemOutput) to type string`

    **Solution:** Use "encoding/json" for converting the data to required type

    ```golang
    json,err := json.Marshal(result)
    stringVal := string(json)
    ```

* **Error:** `json: Unmarshal(non-pointer main.TestItem)`

  Below code doesn't work

  ```go
  insertItem := TestItem{}
  err := json.Unmarshal([]byte(body), insertItem)
  ```

  Correct way to do it

  ```go
  insertItem := &TestItem{}
  err := json.Unmarshal([]byte(body), insertItem)
  ```

* **Error:** `remote: Repository not found. fatal: repository 'https://github.com/aasisodiya/aws/' not found`

  **Solution:** Check mod file it it contains this line or similar `replace github.com/aasisodiya/aws => ../../aws` and make sure it is pointing to correct folder

## Reference

* [AWS DynamoDB Go Lang SDK Documentation](https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/)
* [PutItem API operation for Amazon DynamoDB](https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.PutItem)
* [Amazon DynamoDB Examples Using the AWS SDK for Go](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/using-dynamodb-with-go-sdk.html)
* [AWS Doc DynamoDB SDK Examples](https://github.com/awsdocs/aws-doc-sdk-examples/tree/master/go/example_code/dynamodb)
* [AWS Doc DynamodDB SDK Example - CreateItem](https://github.com/awsdocs/aws-doc-sdk-examples/blob/master/go/example_code/dynamodb/create_item.go)
* [Remove empty value from elements](https://stackoverflow.com/questions/48257658/remove-elements-with-empty-value-when-no-values)
* [Read Item from DynamoDB Table](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/dynamo-example-read-table-item.html)
* [Basic Item Struct doesn't work](https://stackoverflow.com/questions/46688165/cant-unmarshall-dynamodb-attribute)
* [Convert String to JSON / Struct](https://stackoverflow.com/questions/40429296/converting-string-to-json-or-struct)
* [Marshaling & Unmarshaling AttributeValue](https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/dynamodbattribute/)
--
* [AWS DynamoDB Go Lang SDK Documentation - Delete Item from DynamoDB](https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.DeleteItem)
