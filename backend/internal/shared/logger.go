// Package shared provides shared utilities across the application.
package shared

import (
	"log/slog"
	"os"
	"strings"
)

// Logger is the global application logger
var Logger *slog.Logger

// InitLogger initializes the application logger
func InitLogger(env string) {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	var handler slog.Handler
	var format string

	if strings.ToLower(env) == "production" {
		format = os.Getenv("LOG_FORMAT")
		if format == "" {
			format = "json"
		}
	} else {
		format = os.Getenv("LOG_FORMAT")
		if format == "" {
			format = "text"
		}
	}

	if format == "json" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	Logger = slog.New(handler)
	slog.SetDefault(Logger)
}

// Info logs an info message
func Info(msg string, args ...any) {
	Logger.Info(msg, args...)
}

// Debug logs a debug message
func Debug(msg string, args ...any) {
	Logger.Debug(msg, args...)
}

// Warn logs a warning message
func Warn(msg string, args ...any) {
	Logger.Warn(msg, args...)
}

// Error logs an error message
func Error(msg string, args ...any) {
	Logger.Error(msg, args...)
}

// Fatal logs a fatal message and exits
func Fatal(msg string, args ...any) {
	Logger.Error(msg, args...)
	os.Exit(1)
}
