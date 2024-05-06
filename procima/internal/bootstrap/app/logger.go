package app

import (
	"errors"
	loggerPkg "github.com/SShlykov/procima/procima/pkg/logger"
)

func (app *App) initLogger() error {
	if app.config == nil {
		return errors.New("config is nil")
	}
	app.logger = loggerPkg.SetupLogger(app.config.Logger.Level, app.config.Logger.Mode, app.config.AppName, app.config.Host)
	return nil
}
