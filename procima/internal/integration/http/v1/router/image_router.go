package router

import (
	"github.com/SShlykov/procima/procima/internal/integration/http/middleware"
	loggerPkg "github.com/SShlykov/procima/procima/pkg/logger"
	"github.com/SShlykov/procima/procima/pkg/metrics"
	"github.com/gin-gonic/gin"
)

type ImageController interface {
	ProcessImage(c *gin.Context)
}

func ImageRouter(controller ImageController, metr metrics.Metrics) func(engine *gin.Engine, logger loggerPkg.Logger) {
	return func(engine *gin.Engine, logger loggerPkg.Logger) {
		group := engine.Group(BaseURL)
		group.Use(middleware.Metrics(logger, metr))

		group.POST(ImageUploadURL, controller.ProcessImage)
	}
}
