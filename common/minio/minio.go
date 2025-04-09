package utilminio

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"field-service/config"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
	"net/url"
	"strings"
	"time"
)

type MinioClient struct {
	minio *minio.Client
}

type IMinioClient interface {
	UploadFile(context.Context, string, string) (string, error)
}

func NewMinioClient(minio *minio.Client) IMinioClient {
	return &MinioClient{minio: minio}
}

func (m *MinioClient) UploadFile(ctx context.Context, filename, base64Str string) (string, error) {
	if !strings.HasPrefix(base64Str, "data:") {
		return "", errors.New("invalid base64 format")
	}

	parts := strings.SplitN(base64Str, ",", 2)
	if len(parts) != 2 {
		return "", errors.New("invalid base64 data")
	}

	headerParts := strings.SplitN(parts[0], ";", 2)
	if len(headerParts) != 2 {
		return "", errors.New("invalid base64 header")
	}

	contentType := strings.TrimPrefix(headerParts[0], "data:")
	dataBytes, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		logrus.Errorf("failed to decode base64: %v", err)
		return "", err
	}

	// Upload ke MinIO
	reader := bytes.NewReader(dataBytes)
	_, err = m.minio.PutObject(ctx, config.Config.Minio.BucketName, filename, reader, int64(len(dataBytes)), minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		logrus.Errorf("failed to upload to MinIO: %v", err)
		return "", err
	}

	// Buat Presigned URL (berlaku 1 jam) (preview image), kalau ingin get image lagi dari database lewat base64nya
	reqParams := make(url.Values)
	presignedURL, err := m.minio.PresignedGetObject(ctx, config.Config.Minio.BucketName, filename, time.Hour, reqParams)
	if err != nil {
		logrus.Errorf("failed to generate presigned url: %v", err)
		return "", err
	}

	return presignedURL.String(), nil
}
