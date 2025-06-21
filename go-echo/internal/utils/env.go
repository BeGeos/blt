package utils

import (
	"log"
	"os"
	"slices"

	"github.com/BeGeos/go-echo/internal/settings"
)

func GetEnv() string {
	Env := os.Getenv(settings.AppEnvKey)
	if !slices.Contains(settings.AllowedEnvs, Env) {
		log.Fatalf("APP_ENV is not set or invalid: %s", Env)
	}

	return Env
}

var Env = GetEnv()
