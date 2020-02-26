package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handler)
}

// Request Struct
type Request struct {
	ID    float64 `json:"id"`
	Value string  `json:"value"`
}

// Response Struct
type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

// Handler Method
func Handler(request Request) (Response, error) {
	return Response{
		Message: fmt.Sprintf("Processed request ID: %f", request.ID),
		Ok:      true,
	}, nil
}
