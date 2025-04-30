package minio

import (
	"bytes"
	"context"
	"field-service/config"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
)

type MinioClient struct {
	minio *minio.Client
}

type IMinioClient interface {
	UploadFile(context.Context, string, string, []byte) (string, error)
}

func NewMinioClient(minio *minio.Client) IMinioClient {
	return &MinioClient{minio: minio}
}

func (m *MinioClient) UploadFile(ctx context.Context, filename string, contentType string, data []byte) (string, error) {
	// Upload ke MinIO
	reader := bytes.NewReader(data)
	_, err := m.minio.PutObject(ctx, config.Config.Minio.BucketName, filename, reader, int64(len(data)), minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		logrus.Errorf("failed to upload to MinIO: %v", err)
		return "", err
	}

	// Buat public URL
	publicURL := fmt.Sprintf("%s/%s/%s", config.Config.Minio.Address, config.Config.Minio.BucketName, filename)

	return publicURL, nil
}
