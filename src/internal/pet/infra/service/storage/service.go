package storage

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
	"github.com/college-redberry/open-adoption/internal/pet/infra/constants"
)

type GCPStorageService struct {
	client     *storage.Client
	bucketName string
}

func New() *GCPStorageService {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to create GCP storage client: %w", err))
	}

	return &GCPStorageService{
		client:     client,
		bucketName: constants.GCP_BUCKET_NAME,
	}
}

func (s *GCPStorageService) UploadFile(objectName string, data []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	wc := s.client.Bucket(s.bucketName).Object(fmt.Sprintf("%s/%s", constants.GCP_BUCKET_IMAGES_FOLDER, objectName)).NewWriter(ctx)
	if _, err := wc.Write(data); err != nil {
		return fmt.Errorf("failed to write to GCP object: %w", err)
	}

	if err := wc.Close(); err != nil {
		return fmt.Errorf("failed to close GCP object writer: %w", err)
	}

	return nil
}

func (s *GCPStorageService) GenerateUrl(objectName string) (string, error) {
	url, err := storage.SignedURL(s.bucketName, objectName, nil)
	if err != nil {
		return "", fmt.Errorf("failed to generate signed URL: %w", err)
	}

	return url, nil
}
