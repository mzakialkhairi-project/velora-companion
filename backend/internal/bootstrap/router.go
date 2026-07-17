// Package bootstrap provides application startup functionality.
package bootstrap

import (
	"github.com/gin-gonic/gin"

	"github.com/mzakiaklhairi/velora/internal/shared"
)

// Router holds the Gin router
type Router struct {
	Engine *gin.Engine
}

// InitRouter initializes the Gin router
func InitRouter(debug bool) *Router {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	// Recovery middleware
	engine.Use(gin.Recovery())

	shared.Info("Router initialized",
		"mode", gin.Mode(),
	)

	return &Router{Engine: engine}
}

// RegisterRoutes registers all application routes
func (r *Router) RegisterRoutes(
	readyHandler func(*gin.Context),
	healthHandler func(*gin.Context),
	rootHandler func(*gin.Context),
) {
	// Health check endpoints
	r.Engine.GET("/", rootHandler)
	r.Engine.GET("/health", healthHandler)
	r.Engine.GET("/ready", readyHandler)
}
