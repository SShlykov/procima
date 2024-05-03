package middleware

import (
	loggerPkg "github.com/SShlykov/procima/go_pkg/logger"
	"github.com/gin-gonic/gin"
	"time"
)

func Metrics(logger loggerPkg.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t).Truncate(time.Millisecond).String()

		logger.Info("request", loggerPkg.Any("latency", latency), loggerPkg.Any("status", c.Writer.Status()))
	}
}
