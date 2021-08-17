package aws

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	aws_session "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Client interface {
	Upload(filename string, file io.Reader) (string, error)
}

type s3Client struct {
	session *aws_session.Session
	bucket  string
}

func NewS3Client(accessKey, secretKey, region, bucket string) (S3Client, error) {

	session, err := aws_session.NewSession(
		&aws.Config{
			Region:      aws.String(region),
			Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		})

	if err != nil {
		return nil, err
	}

	return &s3Client{session, bucket}, nil
}

func (c *s3Client) Upload(filename string, file io.Reader) (string, error) {
	uploader := s3manager.NewUploader(c.session)

	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(c.bucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(filename),
		Body:   file,
	})

	return output.Location, err
}
