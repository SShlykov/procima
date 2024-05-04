package registry

import (
	"errors"
	loggerPkg "github.com/SShlykov/procima/go_pkg/logger"
	"github.com/SShlykov/procima/procima/internal/config"
	"github.com/SShlykov/procima/procima/internal/domain/services"
	"github.com/SShlykov/procima/procima/internal/integration/http/endpoint"
	cntr "github.com/SShlykov/procima/procima/internal/integration/http/v1/controller"
	"github.com/SShlykov/procima/procima/internal/integration/http/v1/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitWebServer(logger loggerPkg.Logger, configPath string) (*endpoint.WebServer, error) {
	cfg, err := config.LoadServerConfig(configPath)
	if err != nil {
		return nil, errors.New("failed to load server config: " + err.Error())
	}

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())

	SetRouter(engine, logger, cfg)

	srv := &http.Server{
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		ReadTimeout:       cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IddleTimeout,

		Addr:    cfg.Addr,
		Handler: engine,
	}

	return &endpoint.WebServer{Server: srv, Config: cfg}, nil
}

func SetRouter(engine *gin.Engine, logger loggerPkg.Logger, serverConfig *config.ServerConfig) {
	imageController := initImageController(logger, serverConfig)

	routers :=
		[]func(engine *gin.Engine, logger loggerPkg.Logger){
			router.ImageRouter(imageController),
		}

	for _, routes := range routers {
		routes(engine, logger)
	}
}

func initImageController(logger loggerPkg.Logger, serverConfig *config.ServerConfig) cntr.ImageController {
	service := services.NewImageService(logger)
	controller := cntr.NewImageController(service, logger, serverConfig.AvailableTypes, serverConfig.MaxFileSize)
	return controller
}
