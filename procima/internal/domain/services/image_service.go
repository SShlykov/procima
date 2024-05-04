package services

import (
	"context"
	loggerPkg "github.com/SShlykov/procima/go_pkg/logger"
	"github.com/SShlykov/procima/procima/internal/domain/processor"
	"github.com/SShlykov/procima/procima/internal/models"
	_ "image/jpeg"
	"time"
)

// ImageService интерфейс сервиса обработки изображений
type ImageService interface {
	ProcessImage(ctx context.Context, request models.RequestImage) (*[]byte, error)
}

// imageService сервис обработки изображений
type imageService struct {
	processorChan chan<- processor.ImageProcessorItem
	logger        loggerPkg.Logger
}

// NewImageService создание нового сервиса обработки изображений
func NewImageService(logger loggerPkg.Logger, processorChan chan<- processor.ImageProcessorItem) ImageService {
	return &imageService{logger: logger, processorChan: processorChan}
}

// ProcessImage обработка изображения
func (is *imageService) ProcessImage(ctx context.Context, request models.RequestImage) (*[]byte, error) {
	var res processor.ImageResult
	start := time.Now()

	img, err := base64ToImage(request.Image)
	if err != nil {
		is.logger.Error("prepare image", loggerPkg.Err(err))
		return nil, err
	}

	chanel := make(chan processor.ImageResult)
	is.processorChan <- processor.ImageProcessorItem{Img: img, Operation: request.Operation, Channel: chanel}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case res = <-chanel:
		close(chanel)
	}

	is.logger.Debug("operations", loggerPkg.Any("latency", time.Since(start).Truncate(time.Millisecond).String()))

	return res.Res, res.Err
}
