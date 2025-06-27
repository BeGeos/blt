package settings

import (
	"log"
	"os"
	"slices"
)

func GetEnv() string {
	Env := os.Getenv(AppEnvKey)
	if !slices.Contains(AllowedEnvs, Env) {
		log.Fatalf("APP_ENV is not set or invalid: %s", Env)
	}

	return Env
}

var Env = GetEnv()
