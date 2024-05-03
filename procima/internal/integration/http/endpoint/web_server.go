package endpoint

import (
	"context"
	loggerPkg "github.com/SShlykov/procima/go_pkg/logger"
	"github.com/SShlykov/procima/procima/internal/config"
	"net/http"
)

type WebServer struct {
	Server *http.Server
	Config *config.ServerConfig
}

func (web *WebServer) Run(ctx context.Context, logger loggerPkg.Logger) error {
	go func() {
		logger.Info("web server started", loggerPkg.Any("addr", web.Config.Addr))
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
