package integration

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// var AccessKeyID string
// var SecretAccessKey string
// var MyRegion string
var myBucket string

// UploadFile uploads the passed path to cotu bucket and returns
// a path.
func UploadFile(filepath string) (err error) {
	sess := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(sess)

	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open file %q, %v", filepath, err)
	}
	defer file.Close()

	myBucket = getEnvWithKey("AWS_BUCKET_NAME")

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String(filepath),
		Body:   file,
	})

	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	fmt.Printf("file uploaded to: %s\n", aws.StringValue(&result.Location))
	return nil
}

// ListFiles list files from cotu bucket.
func ListFiles() (err error) {
	myBucket = getEnvWithKey("AWS_BUCKET_NAME")
	svc := s3.New(session.New())
	input := &s3.ListObjectsV2Input{
		Bucket:  aws.String(myBucket),
		MaxKeys: aws.Int64(10),
	}

	result, err := svc.ListObjectsV2(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				return aerr
			default:
				// generic errors
				return aerr
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
	return nil
}

// getEnvWithKey gets env value for given key
func getEnvWithKey(key string) string {
	return os.Getenv(key)
}
