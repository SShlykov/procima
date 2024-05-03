package controller

import (
	"context"
	"fmt"
	loggerPkg "github.com/SShlykov/procima/go_pkg/logger"
	"github.com/SShlykov/procima/procima/internal/integration/http/v1/errors"
	"github.com/SShlykov/procima/procima/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ImageController interface {
	ProcessImage(c *gin.Context)
}

type ImageService interface {
	ProcessImage(ctx context.Context, request models.RequestImage) (*models.Image, error)
}

type imageController struct {
	service ImageService
	logger  loggerPkg.Logger
}

func NewImageController(service ImageService, logger loggerPkg.Logger) ImageController {
	return &imageController{service: service, logger: logger}
}

func (ic *imageController) ProcessImage(c *gin.Context) {
	ctx := c.Request.Context()

	var request models.RequestImage
	if err := c.BindJSON(&request); err != nil {
		ic.logger.Error(errors.ErrorBadRequest, loggerPkg.Err(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrorBadRequest})
		return
	}

	image, err := ic.service.ProcessImage(ctx, request)
	if err != nil {
		ic.logger.Error("error", loggerPkg.Err(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrorInternal})
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", image.Name))
	c.Data(200, "application/octet-stream", image.Data)
}
