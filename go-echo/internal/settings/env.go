package settings

import (
	"log"
	"os"
	"slices"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	var envFile string
	switch Env {
	case EnvDevelopment:
		envFile = ".env.local"
	case EnvStaging:
		envFile = ".env.staging"
	case EnvProduction:
		envFile = ".env"
	}

	err := godotenv.Load(envFile)
	return err
}

func GetEnv() string {
	env := os.Getenv(AppEnvKey)
	if !slices.Contains(AllowedEnvs, env) {
		log.Fatalf("APP_ENV is not set or invalid: %s", env)
	}

	return env
}

var Env = GetEnv()
