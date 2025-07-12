package sentry

import (
	"go-hex/internal/config"

	"github.com/getsentry/sentry-go"
)

type Client struct {
	Options sentry.ClientOptions
}

func GetSentryClientOptions(env string) sentry.ClientOptions {
	switch env {
	case config.EnvStaging:
		return getStagingClientOptions()
	case config.EnvProduction:
		return getProductionClientOptions()
	}

	return sentry.ClientOptions{}
}

func getProductionClientOptions() sentry.ClientOptions {
	cfg, _ := config.Load()
	return sentry.ClientOptions{
		Debug:            true,
		AttachStacktrace: true,
		EnableTracing:    true,
		SendDefaultPII:   true,
		TracesSampleRate: 0.3,
		Dsn:              cfg.Sentry.Dsn,
		Environment:      config.EnvProduction,
	}
}

func getStagingClientOptions() sentry.ClientOptions {
	cfg, _ := config.Load()
	return sentry.ClientOptions{
		Debug:            true,
		AttachStacktrace: true,
		TracesSampleRate: 0.8,
		Dsn:              cfg.Sentry.Dsn,
		Environment:      config.EnvStaging,
	}
}

func (c *Client) Init() error {
	return sentry.Init(c.Options)
}
