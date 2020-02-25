package main

import (
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleRequest)
}

// HandleRequest Method
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
