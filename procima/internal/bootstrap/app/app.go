package app

import (
	"context"
	loggerPkg "github.com/SShlykov/procima/go_pkg/logger"
	"github.com/SShlykov/procima/procima/internal/bootstrap/registry"
	configPkg "github.com/SShlykov/procima/procima/internal/config"
	"github.com/SShlykov/procima/procima/internal/domain/processor"
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
	imageProcessorChan := make(chan processor.ImageProcessorItem)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer app.cancel()
		defer app.logger.Info(app.config.AppName + " остановлен")
		server, err := registry.InitWebServer(app.logger, app.configPath, imageProcessorChan)
		if err != nil {
			app.logger.Error("failed to init web server", loggerPkg.Err(err))
			return
		}
		_ = server.Run(ctx, app.logger)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer app.cancel()
		defer app.logger.Info("ImageProcessor остановлен")
		_ = registry.RunImageProcessors(ctx, app.logger, app.configPath, 100, imageProcessorChan)
	}()

	go func() {
		wg.Wait()
		stoppedChan <- struct{}{}
	}()

	return app.closer(ctx, stoppedChan)
}
