package middleware

import (
	loggerPkg "github.com/SShlykov/procima/procima/pkg/logger"
	"github.com/SShlykov/procima/procima/pkg/metrics"
	"github.com/gin-gonic/gin"
	"time"
)

func Metrics(logger loggerPkg.Logger, metr metrics.Metrics) gin.HandlerFunc {
	return func(c *gin.Context) {
		metr.Request()
		t := time.Now()

		c.Next()

		latency := time.Since(t).Truncate(time.Millisecond).String()
		metr.ResponseDuration(float64(time.Since(t).Milliseconds()))

		logger.Info("request", loggerPkg.Any("latency", latency), loggerPkg.Any("status", c.Writer.Status()))
	}
}