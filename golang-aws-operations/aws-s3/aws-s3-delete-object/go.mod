module github.com/aasisodiya/s3

go 1.13

replace github.com/aasisodiya/aws => ../../aws

require (
	github.com/aasisodiya/aws v0.0.0-00010101000000-000000000000
	github.com/aws/aws-lambda-go v1.14.0
	github.com/aws/aws-sdk-go v1.34.0 // indirect
)
