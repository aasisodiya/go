package main

import (
	"fmt"

	"encoding/json"

	dynamodbop "github.com/aasisodiya/aws/dynamodb"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func main() {
	lambda.Start(HandleRequest)

	// // Use below code for local testing
	// request := events.APIGatewayProxyRequest{
	// 	Headers: map[string]string{
	// 		"tableName": "test",
	// 		"region":    "ap-south-1",
	// 	},
	// 	Body:       `{"key1":"value3","pkey":"pkey3"}`,
	// 	HTTPMethod: "GET",
	// }
	// response, _ := HandleRequest(request)
	// fmt.Println("Response from HandleRequest:", response)
}

// TestItem is a struct for your dynamoDB item
type TestItem struct {
	Key1 string `json:"key1"`
	Pkey string `json:"pkey"`
}

// HandleRequest Function
func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod == "PUT" || request.HTTPMethod == "POST" {
		tableName := request.Headers["tableName"]
		region := request.Headers["region"]
		body := request.Body

		// Check if all required values are received in header
		if tableName == "" || region == "" || body == "" {
			APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"Headers Parameters: [tableName, region] and Request Body all are required\"}", StatusCode: 400}
			return APIResponse, nil
		}

		// Build the record for insertion
		insertItem := &TestItem{}
		err := json.Unmarshal([]byte(body), insertItem)
		if err != nil {
			fmt.Println(err)
			APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"Request Body is Invalid\"}", StatusCode: 502}
			return APIResponse, nil
		}

		// Convert the record to map[string]*dynamodb.AttributeValue
		insertRecord, err := dynamodbattribute.MarshalMap(insertItem)
		if err != nil {
			fmt.Println(err)
			APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"Failed to DynamoDB marshal Record\"}", StatusCode: 502}
			return APIResponse, nil
		}

		// Insert Record in table
		fmt.Println("Inserting record into table", tableName, " in Region", region)
		err = dynamodbop.PutItem(region, tableName, insertRecord)

		// Check for errors
		if err != nil {
			fmt.Println(err.Error())
			APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"" + err.Error() + "\"}", StatusCode: 502}
			return APIResponse, nil
		}

		// Return the results (convert json to string)
		APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"Record Inserted\"}", StatusCode: 200, IsBase64Encoded: false}
		return APIResponse, nil
	}

	// Return error if Method Type is different
	APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"Method Not Allowed!\"}", StatusCode: 405}
	return APIResponse, nil
}
