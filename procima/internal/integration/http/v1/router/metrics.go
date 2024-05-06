package router

import (
	loggerPkg "github.com/SShlykov/procima/procima/pkg/logger"
	"github.com/SShlykov/procima/procima/pkg/metrics"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricsRouter(metr metrics.Metrics) func(engine *gin.Engine, logger loggerPkg.Logger) {
	return func(engine *gin.Engine, _ loggerPkg.Logger) {
		group := engine.Group(MetricsURL)

		group.GET("", prometheusHandler(metr))
	}
}

func prometheusHandler(metr metrics.Metrics) gin.HandlerFunc {
	reg := prometheus.NewRegistry()

	metr.Register(reg)

	h := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
