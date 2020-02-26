package main

import (
	"fmt"
	"strings"
	"github.com/aasisodiya/aws/s3"
	"github.com/aws/aws-lambda-go/events"
	// "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// lambda.Start(HandleRequest)
	request := events.APIGatewayProxyRequest{
		Headers: map[string]string{
			"objectkeys":  "ok.txt,ok1.txt",
			"bucketname": "test-bucket-delete-later-21",
			"region":     "ap-south-1",
		},
		HTTPMethod: "GET",
	}
	response, _ := HandleRequest(request)
	fmt.Println(response)
}

// HandleRequest Method (Using Headers)
func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod == "GET" {
		if request.Headers["objectkeys"] == "" || request.Headers["bucketname"] == "" || request.Headers["region"] == "" {
			APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"Headers Parameters: [objectkeys, bucketname, region] all values are required\"}", StatusCode: 400}
			return APIResponse, nil
		}
		fmt.Println("Deleting", request.Headers["objectkeys"], " from Bucket", request.Headers["bucketname"])
		objectkeys := strings.Split(request.Headers["objectkeys"], ",")
		_, err := s3.DeleteObjects(request.Headers["region"], request.Headers["bucketname"], objectkeys)
		if err != nil {
			fmt.Println(err.Error())
			APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"" + err.Error() + "\"}", StatusCode: 502}
			return APIResponse, nil
		}
		APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"" + request.Headers["objectkeys"] + " deleted from bucket " + request.Headers["bucketname"] + "\"}", StatusCode: 200, IsBase64Encoded: true}
		return APIResponse, nil
	}
	APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"Method Not Allowed!\"}", StatusCode: 405}
	return APIResponse, nil
}
