package router

import (
	loggerPkg "github.com/SShlykov/procima/go_pkg/logger"
	"github.com/SShlykov/procima/procima/internal/integration/http/middleware"
	"github.com/gin-gonic/gin"
)

type ImageController interface {
	ProcessImage(c *gin.Context)
}

func ImageRouter(controller ImageController) func(engine *gin.Engine, logger loggerPkg.Logger) {
	return func(engine *gin.Engine, logger loggerPkg.Logger) {
		group := engine.Group(BaseURL)
		group.Use(middleware.Metrics(logger))

		group.POST(ImageUploadURL, controller.ProcessImage)
	}
}
