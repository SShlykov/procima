package services

import (
	"context"
	loggerPkg "github.com/SShlykov/procima/go_pkg/logger"
	"github.com/SShlykov/procima/procima/internal/models"
	"github.com/SShlykov/procima/procima/internal/models/adapters"
	"image"
	_ "image/jpeg"
)

// ImageService интерфейс сервиса обработки изображений
type ImageService interface {
	ProcessImage(ctx context.Context, request models.RequestImage) (*models.Image, error)
}

// imageService сервис обработки изображений
type imageService struct {
	logger loggerPkg.Logger
}

// NewImageService создание нового сервиса обработки изображений
func NewImageService(logger loggerPkg.Logger) ImageService {
	return &imageService{logger: logger}
}

// ProcessImage обработка изображения
func (is *imageService) ProcessImage(ctx context.Context, request models.RequestImage) (*models.Image, error) {
	var img *image.RGBA
	var err error

	if img, err = base64ToImage(request.Image); err != nil {
		return nil, err
	}

	for _, r := range request.Operations {
		switch r {
		case "rotate:90deg":
			img = Parallel(img, 8, Rotate90deg)
		//img = Rotate90deg(img, 0, img.Bounds().Max.X)
		case "sync_rotate:90deg":
			img = Rotate90deg(img, 0, img.Bounds().Max.X)

		//case "recolor:baw": // black and white
		//	img = recolor(img, colorToBAW)
		//case "recolor:negative":
		//	img = recolor(img, colorToNegative)
		default:
			is.logger.Warn("unknown operation", loggerPkg.Any("operation", r))
			return nil, ErrorUnknownOperation
		}
	}

	return adapters.ImageToModel(img)
}
