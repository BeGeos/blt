package echo

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"

	settings "go-hex/internal/config"
)

type MiddlewareConfig struct{}

func (c *MiddlewareConfig) Cors() middleware.CORSConfig {
	cfg, _ := settings.Load()

	allowedOrigins := []string{}
	switch cfg.Env {
	case settings.EnvDevelopment:
		allowedOrigins = []string{"*"}
	case settings.EnvProduction:
		allowedOrigins = []string{"https://your-production-domain.com"}
	case settings.EnvStaging:
		allowedOrigins = []string{"https://your-staging-domain.com"}

	}

	return middleware.CORSConfig{
		AllowOrigins:     allowedOrigins,
		AllowCredentials: true,
		MaxAge:           15, // 15 seconds
	}
}

func (c *MiddlewareConfig) AuthenticationRateLimiter() middleware.RateLimiterConfig {
	return middleware.RateLimiterConfig{
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: rate.Every(time.Minute / 5), Burst: 5, ExpiresIn: 5 * time.Minute},
		),
		DenyHandler: func(c echo.Context, identifier string, err error) error {
			return echo.ErrTooManyRequests
		},
	}
}

func (c *MiddlewareConfig) Logger() middleware.RequestLoggerConfig {
	return middleware.RequestLoggerConfig{
		LogStatus:     true,
		LogLatency:    true,
		LogProtocol:   true,
		LogRequestID:  true,
		LogRemoteIP:   true,
		LogUserAgent:  true,
		LogURI:        true,
		LogMethod:     true,
		LogError:      true,
		HandleError:   true,
		LogValuesFunc: LogValuesFunc,
	}
}

func NewMiddlewareConfig() *MiddlewareConfig {
	return &MiddlewareConfig{}
}
