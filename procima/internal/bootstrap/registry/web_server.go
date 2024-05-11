package registry

import (
	"errors"
	"github.com/SShlykov/procima/procima/internal/app/http/endpoint"
	"github.com/SShlykov/procima/procima/internal/app/http/middleware"
	v1 "github.com/SShlykov/procima/procima/internal/app/http/v1"
	cntr "github.com/SShlykov/procima/procima/internal/app/http/v1/images"
	metrics2 "github.com/SShlykov/procima/procima/internal/app/http/v1/metrics"
	"github.com/SShlykov/procima/procima/internal/config"
	"github.com/SShlykov/procima/procima/internal/domain/services"
	loggerPkg "github.com/SShlykov/procima/procima/pkg/logger"
	"github.com/SShlykov/procima/procima/pkg/metrics"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitWebServer(logger loggerPkg.Logger, configPath string, metr metrics.Metrics,
	imgService services.ImageService) (*endpoint.WebServer, error) {
	cfg, err := config.LoadServerConfig(configPath)
	if err != nil {
		return nil, errors.New("failed to load server config: " + err.Error())
	}

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.Metrics(logger, metr))
	group := engine.Group(v1.BaseURL)

	metrContr := metrics2.NewMetricsController(metr)
	metrContr.RegisterRoutes(group)

	imgCntr := cntr.NewImageController(imgService, logger, cfg.AvailableTypes, cfg.MaxFileSize)
	imgCntr.RegisterRoutes(group)

	webSever := configToWebServer(cfg)
	webSever.Server.Handler = engine

	return webSever, nil
}

func configToWebServer(cfg *config.ServerConfig) *endpoint.WebServer {
	srv := &http.Server{
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		ReadTimeout:       cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IddleTimeout,

		Addr: cfg.Addr,
	}

	return &endpoint.WebServer{Server: srv, Config: cfg}
}
