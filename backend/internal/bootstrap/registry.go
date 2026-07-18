// Package bootstrap provides application startup functionality.
package bootstrap

import (
	"github.com/mzakiaklhairi/velora/internal/ai/provider"
	"github.com/mzakiaklhairi/velora/internal/ai/provider/ollama"
	"github.com/mzakiaklhairi/velora/internal/shared"
)

// Registry holds all application dependencies
type Registry struct {
	Database *Database
	Redis    *Redis
	Router   *Router
	Config   *shared.Config
}

// NewRegistry creates a new registry
func NewRegistry() *Registry {
	return &Registry{}
}

// Build initializes all dependencies
func (r *Registry) Build(cfg *shared.Config) error {
	r.Config = cfg

	// Initialize database
	db, err := InitDatabase(cfg)
	if err != nil {
		return err
	}
	r.Database = db

	// Initialize Redis
	redisClient, err := InitRedis(cfg)
	if err != nil {
		return err
	}
	r.Redis = redisClient

	// Initialize Ollama provider and register
	ollamaCfg := ollama.NewConfig(cfg)
	ollamaProvider := ollama.NewOllamaProvider(ollamaCfg)
	provider.Register("ollama", ollamaProvider)

	// Initialize router
	router := InitRouter(cfg.AppDebug)
	r.Router = router

	return nil
}

// Close closes all dependencies
func (r *Registry) Close() {
	if r.Database != nil {
		_ = r.Database.Close()
	}
	if r.Redis != nil {
		_ = r.Redis.Close()
	}
}
