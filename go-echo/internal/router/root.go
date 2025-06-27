package router

import (
	"github.com/BeGeos/go-echo/internal/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterRootRoutes(e *echo.Echo) {
	e.GET("/", handlers.RootHandler)

	e.GET("/ping", handlers.HealthCheckHandler)
}
