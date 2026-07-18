// Package shared provides shared utilities across the application.
package shared

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// Config holds all application configuration from environment variables
type Config struct {
	// Application
	AppName     string `env:"APP_NAME"`
	AppEnv      string `env:"APP_ENV"`
	AppPort     int    `env:"APP_PORT"`
	AppDebug    bool   `env:"APP_DEBUG"`
	AppTimezone string `env:"APP_TIMEZONE"`

	// Database
	DBHost     string `env:"DB_HOST"`
	DBPort     int    `env:"DB_PORT"`
	DBName     string `env:"DB_NAME"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBSSLMode  string `env:"DB_SSLMODE"`

	// Redis
	RedisHost     string `env:"REDIS_HOST"`
	RedisPort     int    `env:"REDIS_PORT"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	RedisDB       int    `env:"REDIS_DB"`

	// JWT
	JWTSecret              string `env:"JWT_SECRET"`
	JWTIssuer              string `env:"JWT_ISSUER"`
	JWTAccessTokenExpires  string `env:"JWT_ACCESS_TOKEN_EXPIRES"`
	JWTRefreshTokenExpires string `env:"JWT_REFRESH_TOKEN_EXPIRES"`

	// Ollama
	OllamaURL     string `env:"OLLAMA_URL"`
	OllamaModel   string `env:"OLLAMA_MODEL"`
	OllamaTimeout int    `env:"OLLAMA_TIMEOUT"`
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file from parent directory (ignore error if file doesn't exist)
	if err := godotenv.Load("../.env"); err != nil {
		logNoEnv()
	} else {
		logEnvLoaded()
	}

	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return cfg, nil
}

// logEnvLoaded logs when .env file is loaded
func logEnvLoaded() {
	if Logger != nil {
		Logger.Info("Loaded configuration from .env")
	} else {
		slog.Info("Loaded configuration from .env")
	}
}

// logNoEnv logs when no .env file is found
func logNoEnv() {
	if Logger != nil {
		Logger.Info("Using environment variables (no .env file found)")
	} else {
		slog.Info("Using environment variables (no .env file found)")
	}
}

// GetDSN returns PostgreSQL connection string
func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode,
	)
}

// GetRedisAddr returns Redis address
func (c *Config) GetRedisAddr() string {
	return fmt.Sprintf("%s:%d", c.RedisHost, c.RedisPort)
}

// GetServerAddr returns server address
func (c *Config) GetServerAddr() string {
	return fmt.Sprintf(":%d", c.AppPort)
}

// GetStringEnv returns string value or default
func GetStringEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

// GetIntEnv returns int value or default
func GetIntEnv(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		if parsed, err := strconv.Atoi(val); err == nil {
			return parsed
		}
	}
	return defaultVal
}

// GetBoolEnv returns bool value or default
func GetBoolEnv(key string, defaultVal bool) bool {
	if val := os.Getenv(key); val != "" {
		if parsed, err := strconv.ParseBool(val); err == nil {
			return parsed
		}
	}
	return defaultVal
}
