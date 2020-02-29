package main

import (
	"fmt"

	dynamodbop "github.com/aasisodiya/aws/dynamodb"
	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	lambda.Start(HandleRequest)

	// Use below code for local testing
	// request := events.APIGatewayProxyRequest{
	// 	Headers: map[string]string{
	// 		"tableName":  "test",
	// 		"region":     "ap-south-1",
	// 		"primarykey": "pkey1",
	// 	},
	// 	HTTPMethod: "DELETE",
	// }
	// response, _ := HandleRequest(request)
	// fmt.Println("Response from HandleRequest:", response)
}

// HandleRequest Method (Using Headers)
func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod == "DELETE" {
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
		fmt.Println("Deleting record from table", tableName, " in Region", region, " with key:", primarykey)
		err := dynamodbop.DeleteItem(region, tableName, key)

		// Check for errors
		if err != nil {
			fmt.Println(err.Error())
			APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"" + err.Error() + "\"}", StatusCode: 502}
			return APIResponse, nil
		}

		// Return the results (convert json to string)
		APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"Record Deleted\"}", StatusCode: 200, IsBase64Encoded: false}
		return APIResponse, nil
	}

	// Return error if Method Type is different
	APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"Method Not Allowed!\"}", StatusCode: 405}
	return APIResponse, nil
}
