package domain

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/SShlykov/procima/procima/internal/models"
)

type ImageService interface {
	ProcessImage(ctx context.Context, request models.RequestImage) (*models.Image, error)
}

type imageService struct {
}

func NewImageService() ImageService {
	return &imageService{}
}

func (is *imageService) ProcessImage(ctx context.Context, request models.RequestImage) (*models.Image, error) {
	data, err := base64.StdEncoding.DecodeString(request.Image)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 image: %w", err)
	}

	return &models.Image{Data: data, Name: "test.jpeg"}, nil
}
