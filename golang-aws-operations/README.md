# AWS Operations in Go Lang

![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go.golang-aws-operations&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)

Refer to Documentation of golang-aws-lambda folder for lambda deployment with APIGW

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

You will have to use this object quite often if you are going to use API GW

>Respective operations detail are given in their respective directory

[![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)](https://visitorbadge.io/status?path=aasisodiya.go)
