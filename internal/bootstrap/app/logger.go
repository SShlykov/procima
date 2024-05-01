package app

import (
	"errors"
	loggerPkg "github.com/SShlykov/procima/pkg/logger"
)

func (app *App) initLogger() error {
	if app.config == nil {
		return errors.New("config is nil")
	}
	app.logger = loggerPkg.SetupLogger(app.config.Logger.Level)
	return nil
}
