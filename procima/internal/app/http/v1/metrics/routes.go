package metrics

import (
	v1 "github.com/SShlykov/procima/procima/internal/app/http/v1"
	"github.com/gin-gonic/gin"
)

func (mc *metricsController) RegisterRoutes(router *gin.RouterGroup) {
	router.POST(v1.MetricsURL, mc.PrometheusHandler)
}
