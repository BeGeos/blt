package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Env      string
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	URL string
}

var (
	cfg  *Config
	once sync.Once
)

func Load() (*Config, error) {
	once.Do(func() {
		env := getEnv("GO_ENV", "dev")

		_ = godotenv.Load(".env")        // always try base
		_ = godotenv.Load(".env." + env) // e.g., .env.local or .env.dev

		cfg = &Config{
			Env: env,
			Server: ServerConfig{
				Port: getEnv("SERVER_PORT", ":8080"),
			},
			Database: DatabaseConfig{
				URL: getEnv("DATABASE_URL", ""),
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
