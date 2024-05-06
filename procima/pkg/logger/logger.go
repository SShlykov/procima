package logger

import (
	"log/slog"
	"os"
)

type HandlerOptions struct {
	SlogOpts slog.HandlerOptions
	Service  string
	Host     string
}

type Logger interface {
	Warn(msg string, attrs ...any)
	Info(msg string, attrs ...any)
	Debug(msg string, attrs ...any)
	Error(msg string, attrs ...any)
}

type loggerImp struct {
	logger *slog.Logger
}

// SetupLogger creates a new logger with the given options.
//
// level (debug, info, warn, error); mode (pretty, json); service; host;
func SetupLogger(level string, mode string, service string, host string) Logger {
	levels := map[string]slog.Level{
		"debug": slog.LevelDebug,
		"info":  slog.LevelInfo,
		"warn":  slog.LevelWarn,
		"error": slog.LevelError,
	}
	opts := HandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: levels[level],
		},
		Service: service,
		Host:    host,
	}

	var handler slog.Handler
	switch {
	case mode == "pretty":
		handler = NewPrettyHandler(os.Stdout, opts)
	case mode == "json":
		handler = NewJsonHandler(os.Stdout, opts)
	default:
		panic("unknown mode")
	}

	return &loggerImp{logger: slog.New(handler)}
}

func (l *loggerImp) Warn(msg string, attrs ...any) {
	l.logger.Warn(msg, attrs...)
}

func (l *loggerImp) Info(msg string, attrs ...any) {
	l.logger.Info(msg, attrs...)
}

func (l *loggerImp) Debug(msg string, attrs ...any) {
	l.logger.Debug(msg, attrs...)
}

func (l *loggerImp) Error(msg string, attrs ...any) {
	l.logger.Error(msg, attrs...)
}
