package infrastructure

import (
	"log/slog"
)

// AppContext holds application-wide context
type AppContext struct {
	Logger *slog.Logger
	Config *AppConfig
}

// AppConfig holds application configuration
type AppConfig struct {
	Name     string
	Env      string
	Port     string
	Debug    bool
	Timezone string
}

// NewAppContext creates a new app context
func NewAppContext(logger *slog.Logger, config *AppConfig) *AppContext {
	return &AppContext{
		Logger: logger,
		Config: config,
	}
}
