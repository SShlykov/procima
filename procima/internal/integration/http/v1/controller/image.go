package controller

import (
	"context"
	errorsPkg "errors"
	loggerPkg "github.com/SShlykov/procima/go_pkg/logger"
	"github.com/SShlykov/procima/procima/internal/domain/services"
	"github.com/SShlykov/procima/procima/internal/integration/http/v1/errors"
	"github.com/SShlykov/procima/procima/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

// ImageController контроллер для обработки изображений
//
//go:generate mockgen -destination=./mocks/mock_image_controller.go -package=mocks github.com/SShlykov/procima/procima/internal/integration/http/v1/controller ImageController
type ImageController interface {
	ProcessImage(c *gin.Context)
}

// ImageService сервис для обработки изображений
//
//go:generate mockgen -destination=./mocks/mock_image_service.go -package=mocks github.com/SShlykov/procima/procima/internal/integration/http/v1/controller ImageService
type ImageService interface {
	ProcessImage(ctx context.Context, request models.RequestImage) (*[]byte, error)
}

type imageController struct {
	service             ImageService
	logger              loggerPkg.Logger
	availableImageTypes []string
	fileSizeLimit       int
}

func NewImageController(service ImageService, logger loggerPkg.Logger, types []string, fileSizeLimit int) ImageController {
	return &imageController{service: service, logger: logger, availableImageTypes: types, fileSizeLimit: fileSizeLimit}
}

func (ic *imageController) ProcessImage(c *gin.Context) {
	ctx := c.Request.Context()

	var request models.RequestImage
	if err := c.BindJSON(&request); err != nil {
		ic.logger.Error(errors.ErrorBadRequest, loggerPkg.Err(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrorBadRequest})
		return
	}

	if len(request.Image) > ic.fileSizeLimit {
		ic.logger.Error(errors.ErrorBadRequest, loggerPkg.String("error", errors.ErrorExcededFileSize))
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrorExcededFileSize, "limit": ic.fileSizeLimit, "actual": len(request.Image)})
		return
	}

	imageType, found := getImageType(request.Image)
	if !found || !ic.isAvailable(imageType) {
		ic.logger.Error(errors.ErrorBadRequest, loggerPkg.String("error", errors.ErrorInvalidImageType))
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrorInvalidImageType, "available": ic.availableImageTypes, "actual": imageType})
		return
	}

	image, err := ic.service.ProcessImage(ctx, request)
	if err != nil {
		ic.logger.Error("error", loggerPkg.Err(err))
		if errorsPkg.Is(err, services.ErrorUnknownOperation) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrorInternal})
		return
	}

	c.Writer.Header().Set("Content-Type", "image/"+imageType)
	c.Writer.Header().Set("Content-Disposition", "inline")
	_, _ = c.Writer.Write(*image)
}

func (ic *imageController) isAvailable(imageType string) bool {
	for _, t := range ic.availableImageTypes {
		if t == imageType {
			return true
		}
	}
	return false
}

const PartsWhenImage = 2

func getImageType(dataURL string) (string, bool) {
	re := regexp.MustCompile(`^data:image/([^;]+);base64,`)
	matches := re.FindStringSubmatch(dataURL[:50])

	if len(matches) == PartsWhenImage {
		return matches[1], true
	}
	return "", false
}
