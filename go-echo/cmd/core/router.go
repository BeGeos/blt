package core

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/labstack/echo/v4"

	"github.com/BeGeos/go-echo/cmd/apps/auth"
)

func registerRootRoutes(e *echo.Echo) {
	e.GET("/", RootHandler).Name = "home"
	e.GET("/ping", HealthCheckHandler).Name = "ping"
}

func RegisterPprofRoutes(e *echo.Echo) {
	e.GET("/debug/*", echo.WrapHandler(http.DefaultServeMux))
}

func RegisterRoutes(e *echo.Echo) {
	// Register your routes here
	registerRootRoutes(e)

	// 2.0
	auth.RegisterAuthRoutes(e)
}
