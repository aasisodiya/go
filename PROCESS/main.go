package main

import (
    "bytes"
    "log"
    // "net/http"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
	"fmt"
	"./archieve"
)

// TODO fill these in!
const (
    S3_REGION = "us-west-2"
    S3_BUCKET = "delete-this-tempo-bucket"
)


func main() {

	// List of Files to Zip
	files := []string{"/home/ubuntu/compress/test.csv"}
	output := "/home/ubuntu/archieve/donegl4.zip"

	if err := archieve.ZipFiles(output, files); err != nil {
		panic(err)
	}
	fmt.Println("Zipped File:", output)
	output2 := "/home/ubuntu/archieve/donecopy.zip"

	if err := archieve.ZipFiles(output2, files); err != nil {
		panic(err)
	}
	fmt.Println("Zipped File:", output)
	// S3 Code Below
	// Create a single AWS session (we can re use this if we're uploading many files)
	s, err := session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	// Upload
	err = AddFileToS3(s, output)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}


// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
func AddFileToS3(s *session.Session, fileDir string) error {

    // Open the file for use
    file, err := os.Open(fileDir)
    if err != nil {
		fmt.Println(err)
        return err
    }
    defer file.Close()

    // Get file size and read the file content into a buffer
    fileInfo, _ := file.Stat()
    var size int64 = fileInfo.Size()
    buffer := make([]byte, size)
    file.Read(buffer)

    // Config settings: this is where you choose the bucket, filename, content-type etc.
    // of the file you're uploading.
    _, err = s3.New(s).PutObject(&s3.PutObjectInput{
        Bucket:               aws.String(S3_BUCKET),
        Key:                  aws.String(fileDir),
        // ACL:                  aws.String("private"),
        Body:                 bytes.NewReader(buffer),
        // ContentLength:        aws.Int64(size),
        // ContentType:          aws.String(http.DetectContentType(buffer)),
        // ContentDisposition:   aws.String("attachment"),
		// ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("GLACIER"),
	})
	fmt.Println(err)
    return err
}