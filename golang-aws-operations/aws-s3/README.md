# AWS S3 Operations in Go Lang

![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go.golang-aws-operations.aws-s3&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)

Refer below request structure for creating your lambda code with api gateway

## APIGatewayProxyRequest Struct

```golang
type APIGatewayProxyRequest struct {
    Resource                        string                        `json:"resource"` // The resource path defined in API Gateway
    Path                            string                        `json:"path"`     // The url path for the caller
    HTTPMethod                      string                        `json:"httpMethod"`
    Headers                         map[string]string             `json:"headers"`
    MultiValueHeaders               map[string][]string           `json:"multiValueHeaders"`
    QueryStringParameters           map[string]string             `json:"queryStringParameters"`
    MultiValueQueryStringParameters map[string][]string           `json:"multiValueQueryStringParameters"`
    PathParameters                  map[string]string             `json:"pathParameters"`
    StageVariables                  map[string]string             `json:"stageVariables"`
    RequestContext                  APIGatewayProxyRequestContext `json:"requestContext"`
    Body                            string                        `json:"body"`
    IsBase64Encoded                 bool                          `json:"isBase64Encoded,omitempty"`
}
```

## High Level Steps Followed

1. Create mod `go mod init github.com/aasisodiya/s3`
2. Add below line to go.mod, where github.com/aasisodiya/aws is replaced with relative path of aws folder containing the required functions

    ```
    replace github.com/aasisodiya/aws => ../../aws
    ```

3. Call the required method from aws inside main.go

## Troubleshooting

* **Error:** MissingRegion: could not find region configuration

  Below code doesn't work

  ```golang
  svc := s3.New(session.New())
  ```

  Use this instead

  ```golang
  svc := s3.New(session.New(&aws.Config{Region: aws.String     ("REPLACE_WITH_YOUR_REGION_CODE")}))
  ```

## Reference

* [S3 Basic Example](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/s3-example-basic-bucket-operations.html)
* [APIGatewayProxyRequest Structure](https://github.com/aws/aws-lambda-go/blob/v1.14.0/events/apigw.go#L6)
* [Convert io.ReadCloser to a String](https://golangcode.com/convert-io-readcloser-to-a-string/)

[![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)](https://visitorbadge.io/status?path=aasisodiya.go)
