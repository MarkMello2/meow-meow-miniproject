package service

import (
	"context"
	"net/url"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
)

type minioService struct {
	minioClient *minio.Client
}

func NewMinioService(minioClient *minio.Client) MinioService {
	return minioService{minioClient: minioClient}
}

func (m minioService) getUrlImagePath(pathImg string) (string, error) {
	reqParams := make(url.Values)
	myBucket := os.Getenv("MINIO_BUCKET")

	presignedURL, err := m.minioClient.PresignedGetObject(context.Background(), myBucket, pathImg, time.Second*24*60*60, reqParams)
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}
