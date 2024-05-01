package registry

import (
	"context"
	"errors"
	"github.com/SShlykov/procima/internal/config"
	"github.com/SShlykov/procima/internal/integration/http/v1/router"
	loggerPkg "github.com/SShlykov/procima/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WebServer struct {
	server *http.Server
	config *config.ServerConfig
}

func (web *WebServer) Run(ctx context.Context, logger loggerPkg.Logger) error {
	go func() {
		logger.Info("web server started", loggerPkg.Any("addr", web.config.Addr))
		_ = web.server.ListenAndServe()
	}()

	return web.closer(ctx)
}

func (web *WebServer) closer(ctx context.Context) error {
	<-ctx.Done()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), web.config.Timeout)
	defer cancel()

	if err := web.server.Shutdown(timeoutCtx); err != nil {
		return err
	}
	return nil
}

func SetRouter(engine *gin.Engine, logger loggerPkg.Logger, ctx context.Context) {
	controllers :=
		[]func(engine *gin.Engine, logger loggerPkg.Logger, ctx context.Context){
			router.Image,
		}

	for _, controller := range controllers {
		controller(engine, logger, ctx)
	}
}

func InitWebServer(ctx context.Context, logger loggerPkg.Logger, configPath string) (*WebServer, error) {
	cfg, err := config.LoadServerConfig(configPath)
	if err != nil {
		return nil, errors.New("failed to load server config: " + err.Error())
	}

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())

	SetRouter(engine, logger, ctx)

	srv := &http.Server{
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		ReadTimeout:       cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IddleTimeout,

		Addr:    cfg.Addr,
		Handler: engine,
	}

	return &WebServer{server: srv, config: cfg}, nil
}
