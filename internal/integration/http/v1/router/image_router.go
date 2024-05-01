package router

import (
	"context"
	"github.com/SShlykov/procima/internal/integration/http/middleware"
	loggerPkg "github.com/SShlykov/procima/pkg/logger"
	"github.com/gin-gonic/gin"
)

func Image(engine *gin.Engine, logger loggerPkg.Logger, _ context.Context) {
	group := engine.Group(BaseURL)
	group.Use(middleware.Metrics(logger))

	group.POST(ImageUploadURL, func(c *gin.Context) {
		logger.Info("upload image")
		c.String(200, "upload image")
	})
}
