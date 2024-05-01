package app

import (
	"context"
	"github.com/SShlykov/procima/internal/bootstrap/registry"
	configPkg "github.com/SShlykov/procima/internal/config"
	loggerPkg "github.com/SShlykov/procima/pkg/logger"
	"os/signal"
	"sync"
	"syscall"
)

type App struct {
	configPath string

	ctx    context.Context
	cancel context.CancelFunc

	logger loggerPkg.Logger
	config *configPkg.Config
}

func New(configPath string) (*App, error) {
	ctx, cancel := context.WithCancel(context.Background())
	app := &App{ctx: ctx, cancel: cancel, configPath: configPath}

	inits := []func() error{
		app.initConfig,
		app.initLogger,
	}

	for _, init := range inits {
		if err := init(); err != nil {
			return nil, err
		}
	}

	return app, nil
}

func (app *App) Run() error {
	ctx, stop := signal.NotifyContext(app.ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup
	stoppedChan := make(chan struct{})

	app.logger.Info("запуск приложения", loggerPkg.Any("cfg", app.config))
	app.logger.Debug("включены отладочные сообщения")

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer app.cancel()
		defer app.logger.Info(app.config.AppName + " остановлен")
		server, err := registry.InitWebServer(app.ctx, app.logger, app.configPath)
		if err != nil {
			app.logger.Error("failed to init web server", loggerPkg.Err(err))
			return
		}
		_ = server.Run(ctx, app.logger)
	}()

	go func() {
		wg.Wait()
		stoppedChan <- struct{}{}
	}()

	return app.closer(ctx, stoppedChan)
}
