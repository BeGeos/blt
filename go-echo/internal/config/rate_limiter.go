package config

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func GetAuthenticationRateLimiterConfig() middleware.RateLimiterConfig {
	return middleware.RateLimiterConfig{
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: rate.Every(time.Minute / 5), Burst: 5, ExpiresIn: 5 * time.Minute},
		),
		DenyHandler: func(c echo.Context, identifier string, err error) error {
			return echo.ErrTooManyRequests
		},
	}
}
