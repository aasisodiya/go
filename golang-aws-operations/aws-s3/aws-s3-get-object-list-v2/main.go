package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aasisodiya/aws/s3"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// HandleRequest Method (Using Headers)
func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod == "GET" {
		if request.Headers["maxkeys"] == "" || request.Headers["bucketname"] == "" || request.Headers["region"] == "" {
			APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"Headers Parameters: [maxkeys, bucketname, region] all values are required\"}", StatusCode: 400}
			return APIResponse, nil
		}
		fmt.Println("Fetching", request.Headers["maxkeys"], "Objects from Bucket", request.Headers["bucketname"])
		maxkeys, err := strconv.Atoi(request.Headers["maxkeys"])
		if err != nil {
			APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"Invalid query string: [maxkeys]\"}", StatusCode: 400}
			return APIResponse, nil
		}
		result, err := s3.ListObjects(request.Headers["region"], request.Headers["bucketname"], maxkeys)
		if err != nil {
			fmt.Println(err.Error())
			APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"" + err.Error() + "\"}", StatusCode: 502}
			return APIResponse, nil
		}
		jsonResponse, err := json.Marshal(result)
		if err != nil {
			panic(err)
		}

		APIResponse := events.APIGatewayProxyResponse{Body: string(jsonResponse), StatusCode: 200}
		return APIResponse, nil
	}
	APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"Method Not Allowed!\"}", StatusCode: 405}
	return APIResponse, nil
}

func main() {
	lambda.Start(HandleRequest)
	// request := events.APIGatewayProxyRequest {
	// 	HTTPMethod: "GET",
	// 	Headers: map[string]string{
	// 		"bucketname":"test-bucket-delete-later",
	// 		"maxkeys":"5",
	// 		"region":"ap-south-1",
	// 	},
	// }
	// response,_ := HandleRequest(request)
	// fmt.Println(response)
}

// // HandleRequest Method (Using QueryParameters)(Not Recommended)
// func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	if request.HTTPMethod == "GET" {
// 		if request.QueryStringParameters["maxkeys"] == "" || request.QueryStringParameters["bucketname"] == "" {
// 			APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"Query Parameters: [maxkeys, bucketname] both values are required\"}", StatusCode: 400}
// 			return APIResponse, nil
// 		}
// 		fmt.Println("Fetching", request.QueryStringParameters["maxkeys"], "Objects from Bucket", request.QueryStringParameters["bucketname"])
// 		maxkeys, err := strconv.Atoi(request.QueryStringParameters["maxkeys"])
// 		if err != nil {
// 			APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"Invalid query string: [maxkeys]\"}", StatusCode: 400}
// 			return APIResponse, nil
// 		}
// 		result, err := s3.ListObjects("ap-south-1", request.QueryStringParameters["bucketname"], maxkeys)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"" + err.Error() + "\"}", StatusCode: 502}
// 			return APIResponse, nil
// 		}
// 		jsonResponse, err := json.Marshal(result)
// 		if err != nil {
// 			panic(err)
// 		}

// 		APIResponse := events.APIGatewayProxyResponse{Body: string(jsonResponse), StatusCode: 200}
// 		return APIResponse, nil
// 	}
// 	APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"Method Not Allowed!\"}", StatusCode: 405}
// 	return APIResponse, nil
// }
