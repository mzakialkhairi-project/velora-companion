// Package bootstrap provides application startup functionality.
package bootstrap

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/mzakiaklhairi/velora/internal/shared"
)

// Redis holds the Redis client
type Redis struct {
	Client *redis.Client
}

// InitRedis initializes the Redis connection
func InitRedis(cfg *shared.Config) (*Redis, error) {
	shared.Info("Connecting to Redis...")

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.GetRedisAddr(),
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		shared.Error("Failed to connect to Redis", "error", err.Error())
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	shared.Info("Redis connected successfully",
		"addr", cfg.GetRedisAddr(),
	)

	return &Redis{Client: client}, nil
}

// Close closes the Redis connection
func (r *Redis) Close() error {
	return r.Client.Close()
}

// Ping checks if Redis is reachable
func (r *Redis) Ping(ctx context.Context) error {
	return r.Client.Ping(ctx).Err()
}
