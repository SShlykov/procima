package metrics

import (
	"github.com/SShlykov/procima/procima/pkg/metrics"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Controller interface {
	PrometheusHandler(c *gin.Context)
	RegisterRoutes(router *gin.RouterGroup)
}

type metricsController struct {
	metr metrics.Metrics
}

func NewMetricsController(metr metrics.Metrics) Controller {
	return &metricsController{metr: metr}
}

func (mc *metricsController) PrometheusHandler(c *gin.Context) {
	reg := prometheus.NewRegistry()

	mc.metr.Register(reg)

	h := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})

	h.ServeHTTP(c.Writer, c.Request)
}
