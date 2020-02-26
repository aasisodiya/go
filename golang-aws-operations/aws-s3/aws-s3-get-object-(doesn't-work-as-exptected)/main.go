package main

import (
	// "encoding/json"
	"fmt"
    "bytes"
	"encoding/base64"

	"github.com/aasisodiya/aws/s3"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// HandleRequest Method
func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println(request)
	fmt.Println(request.Headers)
	if request.HTTPMethod == "GET" {
		if request.Headers["objectkey"] == "" || request.Headers["bucketname"] == "" || request.Headers["region"] == ""{
			APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"Headers Parameters: [objectkey, bucketname, region] all values are required\"}", StatusCode: 400}
			return APIResponse, nil
		}
		fmt.Println("Fetching", request.Headers["objectkey"], " from Bucket", request.Headers["bucketname"])
		result, err := s3.GetObject(request.Headers["region"], request.Headers["bucketname"], request.Headers["objectkey"])
		if err != nil {
			fmt.Println(err.Error())
			APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"" + err.Error() + "\"}", StatusCode: 502}
			return APIResponse, nil
		}
		fmt.Print(result)
		buf := new(bytes.Buffer)
		buf.ReadFrom(result.Body)
		newStr := buf.String()
		encodedString := base64.StdEncoding.EncodeToString([]byte(newStr))
		APIResponse := events.APIGatewayProxyResponse{Body: encodedString, StatusCode: 200, IsBase64Encoded: true}
		return APIResponse, nil
	}
	APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"Method Not Allowed!\"}", StatusCode: 405}
	return APIResponse, nil
}

func main() {
	// request := events.APIGatewayProxyRequest {
	// 	Headers: map[string]string {
	// 		"objectkey":"ok.txt",
	// 		"bucketname":"test-bucket-delete-later-21",
	// 		"region":"ap-south-1",
	// 	},
	// 	HTTPMethod: "GET",
	// }
	// response,_ := HandleRequest(request)
	// fmt.Println(response)
	lambda.Start(HandleRequest)
}