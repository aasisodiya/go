package main

import(
	"fmt"
	"github.com/aasisodiya/aws/s3"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleRequest)
	// request := events.APIGatewayProxyRequest {
	// 	Headers: map[string]string {
	// 		"bucketname":"test-bucket-delete-later-2",
	// 		"region":"ap-south-1",
	// 	},
	// 	HTTPMethod: "GET",
	// }
	// response,_ := HandleRequest(request)
	// fmt.Println(response)
}

// HandleRequest Method (Using Headers)
func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod == "GET" {
		if request.Headers["bucketname"] == "" || request.Headers["region"] == ""{
			APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"Headers Parameters: [bucketname, region] both values are required\"}", StatusCode: 400}
			return APIResponse, nil
		}
		fmt.Println("Creating", request.Headers["bucketname"], " in Region", request.Headers["region"])
		result, err := s3.CreateBucket(request.Headers["region"], request.Headers["bucketname"])
		if err != nil {
			fmt.Println(err.Error())
			APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"" + err.Error() + "\"}", StatusCode: 502}
			return APIResponse, nil
		}
		fmt.Print(result)
		APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \""+"Bucket " +request.Headers["bucketname"] + " created in region" + request.Headers["region"]+"\"}", StatusCode: 200, IsBase64Encoded: true}
		return APIResponse, nil
	}
	APIResponse := events.APIGatewayProxyResponse{Body: "{\"message\": \"Method Not Allowed!\"}", StatusCode: 405}
	return APIResponse, nil
}