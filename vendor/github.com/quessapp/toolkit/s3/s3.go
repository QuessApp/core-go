package s3

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3Credentials represents the AWS S3 credentials.
type S3Credentials struct {
	AccessKey string
	Secret    string
	Token     string
}

// Configure configures AWS S3 client.
func Configure(region *string, S3credentials *S3Credentials) (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      region,
		Credentials: credentials.NewStaticCredentials(S3credentials.AccessKey, S3credentials.Secret, S3credentials.Token),
	})

	if err != nil {
		return nil, err
	}

	return s3.New(sess), nil
}

// UploadFile uploads file to S3.
func UploadFile(client *s3.S3, bucketName, fileName string, file io.ReadSeeker, acl *string) (*s3.PutObjectOutput, error) {
	return client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
		ACL:    acl,
	})
}

// DeleteFile deletes file from S3.
func DeleteFile(client *s3.S3, bucketName, fileName string) (*s3.DeleteObjectOutput, error) {
	return client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
	})
}
