package app

import (
	"context"
	"time"
)

const (
	TerminateTimeout = 10 * time.Second
)

func (app *App) closer(ctx context.Context, stoppedChan <-chan struct{}) error {
	<-ctx.Done()

	app.logger.Info("Остановка приложения")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), TerminateTimeout)
	defer cancel()

	select {
	case <-stoppedChan:
		app.logger.Info("Приложение остановлено")
		return nil
	case <-timeoutCtx.Done():
		app.logger.Error("Приложение остановлено по таймауту")
		return ctx.Err()
	}
}
