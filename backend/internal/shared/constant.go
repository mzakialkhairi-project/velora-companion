// Package shared provides shared utilities across the application.
package shared

// App constants
const (
	AppName    = "Velora"
	AppVersion = "1.0.0"
	AppEnv     = "development"
)

// Default values
const (
	DefaultAppPort     = 18080
	DefaultAppDebug    = true
	DefaultAppTimezone = "Asia/Jakarta"

	DefaultDBHost     = "localhost"
	DefaultDBPort     = 15432
	DefaultDBName     = "velora"
	DefaultDBUser     = "postgres"
	DefaultDBPassword = "postgres"
	DefaultDBSSLMode  = "disable"

	DefaultRedisHost = "localhost"
	DefaultRedisPort = 16379
	DefaultRedisDB   = 0
)
