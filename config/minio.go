package config

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinio() (*minio.Client, error) {
	config := Config
	minioClient, err := minio.New(config.Minio.Address, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Minio.AccessKey, config.Minio.Secret, ""),
		Secure: config.Minio.UseSsl,
	})

	if err != nil {
		return nil, err
	}

	_, err = minioClient.ListBuckets(context.Background())
	if err != nil {
		return nil, err
	}

	return minioClient, nil

}
