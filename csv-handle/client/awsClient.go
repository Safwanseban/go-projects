package client

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AwsS3Client struct {
	svc *s3.S3
}

func NewAwsClient() *AwsS3Client {

	region := "ap-northeast-1"
	sess := session.Must(session.NewSession(&aws.Config{Region: &region}))
	return &AwsS3Client{
		svc: s3.New(sess),
	}
}

func (awsC *AwsS3Client) UploadToS3(bucket, key string, data io.ReadSeeker) error {

	_, err := awsC.svc.PutObject(&s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   data,
	})
	if err != nil {

		return err
	}
	return nil

}
