package middleware

import (
	loggerPkg "github.com/SShlykov/procima/pkg/logger"
	"github.com/gin-gonic/gin"
	"time"
)

func Metrics(logger loggerPkg.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)
		logger.Info("latency", loggerPkg.Any("latency", latency), loggerPkg.Any("status", c.Writer.Status()))
	}
}
