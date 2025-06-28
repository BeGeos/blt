package config

import (
	"github.com/labstack/echo/v4/middleware"

	"github.com/BeGeos/go-echo/internal/settings"
)

func getAllowedOrigins() []string {
	switch settings.Env {
	case "dev":
		return []string{"*"}
	case "prod":
		return []string{"https://your-production-domain.com"}
	case "staging":
		return []string{"https://your-staging-domain.com"}
	}

	return []string{}
}

var CorsConfig = middleware.CORSConfig{
	AllowOrigins:     getAllowedOrigins(),
	AllowCredentials: true,
	MaxAge:           15, // 15 seconds
}
