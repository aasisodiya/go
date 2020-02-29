package s3

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// ListObjects Method
func ListObjects(region string, bucketName string, maxkeys int) (*s3.ListObjectsV2Output, error) {
	fmt.Println("ListObjects from Region:", region, "& BucketName:", bucketName)
	svc := s3.New(session.New(&aws.Config{
		Region: aws.String(region)}))
	input := &s3.ListObjectsV2Input{
		Bucket:  aws.String(bucketName),
		MaxKeys: aws.Int64(int64(maxkeys)),
	}

	result, err := svc.ListObjectsV2(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		}
		return nil, err
	}
	return result, nil
}

// GetObject Method
func GetObject(region string, bucketName string, objectkey string) (*s3.GetObjectOutput, error) {
	fmt.Println("Get Object from Region:", region, "& BucketName:", bucketName, "with Key:", objectkey)
	svc := s3.New(session.New(&aws.Config{
		Region: aws.String(region)}))
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectkey),
	}

	result, err := svc.GetObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				fmt.Println(s3.ErrCodeNoSuchKey, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		}
		return nil, err
	}

	return result, nil
}

// DeleteObject Method
func DeleteObject(region string, bucketName string, objectkey string) (*s3.DeleteObjectOutput, error) {
	fmt.Println("Deleting Object from Region:", region, "& BucketName:", bucketName, "with Key:", objectkey)
	svc := s3.New(session.New(&aws.Config{
		Region: aws.String(region)}))
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectkey),
	}

	result, err := svc.DeleteObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		}
		return nil, err
	}

	return result, nil
}

// DeleteObjects Method
func DeleteObjects(region string, bucketName string, objectkeys []string) (*s3.DeleteObjectsOutput, error) {
	fmt.Println("Deleting Objects from Region:", region, "& BucketName:", bucketName, "with Keys:", objectkeys)
	objects := make([]*s3.ObjectIdentifier, len(objectkeys))
	for i := 0; i < len(objectkeys); i++ {
		objects[i] = &s3.ObjectIdentifier{
			Key: aws.String(objectkeys[i]),
		}
	}

	svc := s3.New(session.New(&aws.Config{
		Region: aws.String(region)}))
	input := &s3.DeleteObjectsInput{
		Bucket: aws.String(bucketName),
		Delete: &s3.Delete{
			Objects: objects,
			Quiet:   aws.Bool(false),
		},
	}

	result, err := svc.DeleteObjects(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		}
		return nil, err
	}

	return result, nil
}

// CreateBucket Method
func CreateBucket(region string, bucketName string) (*s3.CreateBucketOutput, error) {
	fmt.Println("Creating Bucket in region:", region, "& BucketName:", bucketName)
	svc := s3.New(session.New(&aws.Config{
		Region: aws.String(region)}))
	input := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String(region),
		},
	}

	result, err := svc.CreateBucket(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				fmt.Println(s3.ErrCodeBucketAlreadyExists, aerr.Error())
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				fmt.Println(s3.ErrCodeBucketAlreadyOwnedByYou, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		}
		return nil, err
	}

	return result, nil
}

// DeleteBucket Method
func DeleteBucket(region string, bucketName string) (*s3.DeleteBucketOutput, error) {
	fmt.Println("Deleting Bucket in region:", region, "& BucketName:", bucketName)
	svc := s3.New(session.New(&aws.Config{
		Region: aws.String(region)}))
	input := &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	}

	result, err := svc.DeleteBucket(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		}
		return nil, err
	}

	return result, nil
}
