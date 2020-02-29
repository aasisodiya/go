# AWS Operations in Go Lang

Refer to Documentation of golang-aws-lambda folder for lambda deployment with apigw

## APIGatewayProxyRequest Struct

```json
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
