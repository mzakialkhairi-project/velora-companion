// Package main provides the entry point for the Velora API server.
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mzakiaklhairi/velora/internal/bootstrap"
	"github.com/mzakiaklhairi/velora/internal/infrastructure/jwt"
	authrepo "github.com/mzakiaklhairi/velora/internal/modules/auth/repository"
	userrepo "github.com/mzakiaklhairi/velora/internal/modules/user/repository"
	"github.com/mzakiaklhairi/velora/internal/shared"
)

func main() {
	// Load configuration
	cfg, err := shared.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger
	shared.InitLogger(cfg.AppEnv)

	shared.Info("Starting Velora API",
		"app", cfg.AppName,
		"env", cfg.AppEnv,
		"port", cfg.AppPort,
	)

	// Build registry
	registry := bootstrap.NewRegistry()
	if err := registry.Build(cfg); err != nil {
		shared.Fatal("Failed to build application", "error", err.Error())
	}
	defer registry.Close()

	// Create user repository
	userRepo := userrepo.NewPostgresUserRepository(registry.Database.DB)

	// Create refresh token repository
	refreshTokenRepo := authrepo.NewPostgresRefreshTokenRepository(registry.Database.DB)

	// Initialize JWT service
	jwtService, err := jwt.NewJWTService(jwt.Config{
		Secret:        cfg.JWTSecret,
		Issuer:        cfg.JWTIssuer,
		AccessExpires: cfg.JWTAccessTokenExpires,
	})
	if err != nil {
		shared.Fatal("Failed to create JWT service", "error", err.Error())
	}

	// Parse refresh token expiry
	refreshExpires, err := time.ParseDuration(cfg.JWTRefreshTokenExpires)
	if err != nil {
		shared.Fatal("Failed to parse refresh token expiry", "error", err.Error())
	}

	// Register health endpoints
	registry.Router.RegisterRoutes(
		readyHandler(registry),
		healthHandler,
		rootHandler(cfg),
		userRepo,
		jwtService,
		refreshTokenRepo,
		registry.Database.DB,
		refreshExpires,
	)

	// Create HTTP server
	srv := &http.Server{
		Addr:    cfg.GetServerAddr(),
		Handler: registry.Router.Engine,
	}

	// Start server in goroutine
	go func() {
		shared.Info("HTTP server listening", "port", cfg.AppPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			shared.Fatal("Failed to start server", "error", err.Error())
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	shared.Info("Shutting down server...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		shared.Error("Server forced to shutdown", "error", err.Error())
	}

	shared.Info("Server exited")
}

func rootHandler(cfg *shared.Config) func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, shared.Response{
			Success: true,
			Data: map[string]string{
				"name":    cfg.AppName,
				"version": shared.AppVersion,
				"env":     cfg.AppEnv,
			},
		})
	}
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, shared.Response{
		Success: true,
		Data: map[string]string{
			"status": "ok",
		},
	})
}

func readyHandler(registry *bootstrap.Registry) func(*gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()

		// Check PostgreSQL
		if err := registry.Database.Ping(ctx); err != nil {
			c.JSON(http.StatusServiceUnavailable, shared.Response{
				Success: false,
				Error: &shared.ErrorInfo{
					Code:    shared.ErrCodeInternal,
					Message: "PostgreSQL not ready",
				},
			})
			return
		}

		// Check Redis
		if err := registry.Redis.Ping(ctx); err != nil {
			c.JSON(http.StatusServiceUnavailable, shared.Response{
				Success: false,
				Error: &shared.ErrorInfo{
					Code:    shared.ErrCodeInternal,
					Message: "Redis not ready",
				},
			})
			return
		}

		c.JSON(http.StatusOK, shared.Response{
			Success: true,
			Data: map[string]string{
				"status": "ready",
			},
		})
	}
}
