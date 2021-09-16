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
