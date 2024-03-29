package s3

import (
	"context"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/aws/aws-sdk-go-v2/config"
)

type Service struct {
	bucket   string
	prefix   string
	endpoint string
	client   *s3.Client
}

func NewService() Service {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		panic(err)
	}

	bucket := os.Getenv("AWS_S3_BUCKET_NAME")
	if bucket == "" {
		panic("AWS_S3_BUCKET_NAME is required")
	}

	prefix := strings.Trim(os.Getenv("AWS_S3_PREFIX"), "/") + "/"
	if prefix == "/" {
		prefix = ""
	}

	endpoint := os.Getenv("AWS_S3_ENDPOINT")
	if endpoint == "" {
		panic("AWS_S3_ENDPOINT is required")
	}

	client := s3.NewFromConfig(cfg)

	return Service{
		bucket:   bucket,
		prefix:   prefix,
		endpoint: endpoint,
		client:   client,
	}
}

func (s Service) UploadFile(ctx context.Context, key string, body io.Reader) (string, error) {
	key = s.key(key)
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &s.bucket,
		Key:    &key,
		Body:   body,
	})
	if err != nil {
		return "", err
	}

	return s.getFileURL(key), nil
}

func (s Service) key(key string) string {
	return s.prefix + key
}

func (s Service) getFileURL(key string) string {
	dir := filepath.Dir(key)
	fileName := filepath.Base(key)

	return s.endpoint + "/" + filepath.Join(dir, url.PathEscape(fileName))
}

func (s Service) DeleteFile(ctx context.Context, key string) error {
	key = s.key(key)
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &s.bucket,
		Key:    &key,
	})
	return err
}
