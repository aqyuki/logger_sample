package logger

import "log/slog"

var (
	logger *slog.Logger = nil
)

func Info(msg string) {
	logger.Info(msg)
}

func Error(msg string) {
	logger.Error(msg)
}

// etc..

func init() {
	// Any process to initialize logger to use safe
	logger = new(slog.Logger)
}
