package endpoint

import (
	"context"
	"github.com/SShlykov/procima/procima/internal/config"
	loggerPkg "github.com/SShlykov/procima/procima/pkg/logger"
	"net/http"
)

type WebServer struct {
	Server *http.Server
	Config *config.ServerConfig
}

func (web *WebServer) Run(ctx context.Context, logger loggerPkg.Logger) error {
	go func() {
		logger.Info("Веб сервер запущен", loggerPkg.Any("config", web.Config))
		_ = web.Server.ListenAndServe()
	}()

	return web.closer(ctx)
}

func (web *WebServer) closer(ctx context.Context) error {
	<-ctx.Done()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), web.Config.Timeout)
	defer cancel()

	if err := web.Server.Shutdown(timeoutCtx); err != nil {
		return err
	}
	return nil
}
