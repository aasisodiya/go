# Using Parameter Store with Go Lang

![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go.golang-aws-operations.aws-parameter-store&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)

## Create a Parameter in Parameter Store

> Open AWS Parameter Store: [Link](https://us-west-2.console.aws.amazon.com/systems-manager/parameters?region=us-west-2)

Create Parameter that you require. Below names are used for creating parameters for secretkey String

```text
/testapp/secretkey
```

## Define The Same Parameter Name Defined above as Constants preferably in constant.go file

```go
const (
  // SecretKey Parameter name
  SecretKeyParameterName = "/testapp/secretkey"
  // Region
  Region = "us-west-2"
)
```

## Get the Values from Parameter Store

```go
package main

import (
  "log"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/ssm"
  "github.com/pkg/errors"
)

const (
  // SecretKey Parameter name
  SecretKeyParameterName = "/testapp/secretkey"
  // Region
  Region = "us-west-2"
)

func main() {
  val, err := getValueFromParameterStore(SecretKeyParameterName)
  if err != nil {
    log.Println(err)
  }
  log.Println(val, err)
}

func getValueFromParameterStore(ParameterName string) (value string, err error) {
  sess, err := session.NewSession(&aws.Config{
    Region: aws.String(Region)},
  )
  if err != nil {
    return value, errors.Wrap(err, "not found: SecretKeyParameterName")
  }
  ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion(Region))
  param, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
    Name:           aws.String(ParameterName),
    WithDecryption: aws.Bool(false),
  })
  if err != nil {
    return value, errors.Wrap(err, "SecretKeyParameterName")
  }

  value = *param.Parameter.Value
  return value, err
}
```
