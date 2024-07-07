package log

import (
	"fmt"
	"log/slog"
)

// Debug calls [Logger.Debug] on the default logger.
func Debugf(msg string, args ...any) {
	slog.Debug(fmt.Sprintf(msg, args))
}

// Info calls [Logger.Info] on the default logger.
func Infof(msg string, args ...any) {
	slog.Info(fmt.Sprintf(msg, args))
}

// Warn calls [Logger.Warn] on the default logger.
func Warnf(msg string, args ...any) {
	slog.Warn(fmt.Sprintf(msg, args))
}

// Error calls [Logger.Error] on the default logger.
func Errorf(msg string, args ...any) {
	slog.Error(fmt.Sprintf(msg, args))
}
