package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Server         ServerConfig
	Database       DatabaseConfig
	Authentication AuthenticationConfig
	Env            string
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	URL string
}

type AuthenticationConfig struct {
	SigningKey string
	ContextKey string
}

var (
	cfg  *Config
	once sync.Once
)

var (
	EnvDevelopment = "dev"
	EnvStaging     = "staging"
	EnvProduction  = "prod"
)

func Load() (*Config, error) {
	once.Do(func() {
		env := getEnv("GO_ENV", "dev")

		_ = godotenv.Load(".env." + env) // e.g., .env.local or .env.dev
		_ = godotenv.Load(".env")        // from docs - Existing envs take precedence of envs that are loaded later

		cfg = &Config{
			Env: env,
			Server: ServerConfig{
				Port: getEnv("SERVER_PORT", ":8080"),
			},
			Database: DatabaseConfig{
				URL: getEnv("DATABASE_URL", ""),
			},
			Authentication: AuthenticationConfig{
				SigningKey: getEnv("JWT_SIGNING_KEY", ""),
				ContextKey: getEnv("JWT_CONTEXT_KEY", "jwt:token"),
			},
		}
	})

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
