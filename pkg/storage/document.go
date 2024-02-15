package storage

import (
	"context"
	cfg "github.com/VATUSA/primary-api/pkg/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"path"
)

var PublicBucket *StorageClient

type StorageClient struct {
	client *s3.Client
	bucket string
}

func NewS3Client(cfg *cfg.S3Config) (*StorageClient, error) {
	credentialsProvider := credentials.NewStaticCredentialsProvider(cfg.AccessKey, cfg.SecretKey, "")
	s3Config := aws.Config{
		Credentials:  credentialsProvider,
		Region:       cfg.Region,
		BaseEndpoint: aws.String(cfg.Endpoint),
	}

	client := &StorageClient{
		client: s3.NewFromConfig(s3Config, func(options *s3.Options) {}),
		bucket: cfg.Bucket,
	}

	return client, nil
}

func (s *StorageClient) Upload(directory string, filename string, body io.Reader) error {
	fullKey := path.Join(directory, filename)
	_, err := s.client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(fullKey),
		Body:   body,
	})
	return err
}

func (s *StorageClient) Replace(directory string, filename string, body io.Reader) error {
	// In S3, replace is the same as upload. It will overwrite the existing object.
	return s.Upload(directory, filename, body)
}

func (s *StorageClient) Delete(directory, filename string) error {
	fullKey := path.Join(directory, filename)
	_, err := s.client.DeleteObject(context.Background(), &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(fullKey),
	})
	return err
}
