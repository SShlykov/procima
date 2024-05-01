package app

import (
	"context"
	"time"
)

func (app *App) closer(ctx context.Context, stoppedChan <-chan struct{}) error {
	<-ctx.Done()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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
