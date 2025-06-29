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

var AllowedEnvs = []string{"dev", "prod", "staging"}

// Authentication settings

var JwtSecretKey = getEnvKey("JWT_SECRET_KEY", "")
