package main

import (
	"fmt"

	"encoding/json"

	dynamodbop "github.com/aasisodiya/aws/dynamodb"
	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// TestItem is a struct for your dynamoDB item
type TestItem struct {
	Key1 string `json:"key1"`
	Pkey string `json:"pkey"`
}

func main() {
	lambda.Start(HandleRequest)

	// // Use below code for local testing
	// request := events.APIGatewayProxyRequest{
	// 	Headers: map[string]string{
	// 		"tableName":  "test",
	// 		"region":     "ap-south-1",
	// 		"primarykey": "pkey1",
	// 	},
	// 	HTTPMethod: "GET",
	// }
	// response, _ := HandleRequest(request)
	// fmt.Println("Response from HandleRequest:", response)
}

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
		fmt.Println("Result from DynamoDB:", result)

		// Unmarshal the result to a struct
		item := TestItem{}

		err = dynamodbattribute.UnmarshalMap(result.Item, &item)
		if err != nil {
			panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
		}

		// Covert the result to JSON
		json, err := json.Marshal(item)

		// Return the results (convert json to string)
		APIResponse := events.APIGatewayProxyResponse{Body: string(json), StatusCode: 200, IsBase64Encoded: false}
		return APIResponse, nil
	}

	// Return error if Method Type is different
	APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"Method Not Allowed!\"}", StatusCode: 405}
	return APIResponse, nil
}
