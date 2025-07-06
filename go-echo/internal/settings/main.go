package settings

import (
	"os"
)

func getEnvKey(key, defaultValue string) string {
	err := LoadEnv()
	if err != nil {
		return defaultValue
	}

	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultValue
}

const AppEnvKey = "APP_ENV"

const (
	EnvDevelopment = "dev"
	EnvStaging     = "staging"
	EnvProduction  = "prod"
)

var AllowedEnvs = []string{EnvDevelopment, EnvStaging, EnvProduction}

// Authentication settings

var (
	JwtSecretKey  = getEnvKey("JWT_SECRET_KEY", "")
	JwtContextKey = getEnvKey("JWT_CONTEXT_KEY", "tkn")
)

// Sentry
var SentryDsn = getEnvKey("SENTRY_DSN", "")
