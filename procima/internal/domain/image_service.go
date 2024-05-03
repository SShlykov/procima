package domain

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	loggerPkg "github.com/SShlykov/procima/go_pkg/logger"
	"github.com/SShlykov/procima/procima/internal/models"
	"github.com/SShlykov/procima/procima/internal/models/adapters"
	"image"
	_ "image/jpeg"
	"strings"
)

type ImageService interface {
	ProcessImage(ctx context.Context, request models.RequestImage) (*models.Image, error)
}

type imageService struct {
	logger loggerPkg.Logger
}

func NewImageService(logger loggerPkg.Logger) ImageService {
	return &imageService{logger: logger}
}

func (is *imageService) ProcessImage(ctx context.Context, request models.RequestImage) (*models.Image, error) {
	var img image.Image
	var err error

	if img, err = base64ToImage(request.Image); err != nil {
		return nil, err
	}

	for _, r := range request.Operations {
		switch r {
		case "rotate90deg":
			img = rotate90deg(img)
		default:
			is.logger.Warn("unknown operation", loggerPkg.Any("operation", r))
			return nil, errors.New("unknown operation")
		}
	}

	return adapters.ImageToModel(img)
}

func rotate90deg(img image.Image) image.Image {
	bounds := img.Bounds()
	maxX, maxY := bounds.Max.X, bounds.Max.Y
	newImage := image.NewRGBA(image.Rect(0, 0, maxY, maxX))

	for y := bounds.Min.Y; y < maxY; y++ {
		for x := bounds.Min.X; x < maxX; x++ {
			newX := maxY - y - 1
			newY := x
			newImage.Set(newX, newY, img.At(x, y))
		}
	}

	return newImage
}

func base64ToImage(encodedImage string) (image.Image, error) {
	base64String := getBase64String(encodedImage)

	data, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return img, nil
}

func getBase64String(image string) string {
	dataURI := "base64,"
	if idx := strings.Index(image, dataURI); idx > -1 {
		return image[idx+len(dataURI):]
	}
	return ""
}
