package sentry

import (
	"github.com/BeGeos/go-echo/internal/settings"
	"github.com/getsentry/sentry-go"
)

type Client struct {
	Options sentry.ClientOptions
}

func GetSentryClientOptions(env string) sentry.ClientOptions {
	switch env {
	case settings.EnvStaging:
		return getStagingClientOptions()
	case settings.EnvProduction:
		return getProductionClientOptions()
	}

	return sentry.ClientOptions{}
}

func getProductionClientOptions() sentry.ClientOptions {
	return sentry.ClientOptions{
		Debug:            true,
		AttachStacktrace: true,
		EnableTracing:    true,
		SendDefaultPII:   true,
		TracesSampleRate: 0.3,
		Dsn:              settings.SentryDsn,
		Environment:      settings.EnvProduction,
	}
}

func getStagingClientOptions() sentry.ClientOptions {
	return sentry.ClientOptions{
		Debug:            true,
		AttachStacktrace: true,
		TracesSampleRate: 0.8,
		Dsn:              settings.SentryDsn,
		Environment:      settings.EnvStaging,
	}
}

func (c *Client) Init() error {
	return sentry.Init(c.Options)
}
