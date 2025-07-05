package core

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/labstack/echo/v4"

	"github.com/BeGeos/go-echo/cmd/apps/auth"
)

type _Router struct{}

func (r *_Router) registerRootRoutes(e *echo.Echo) {
	e.GET("/", Handlers.RootHandler).Name = "home"
	e.GET("/ping", Handlers.HealthCheckHandler).Name = "ping"
}

func (r *_Router) RegisterPprofRoutes(e *echo.Echo) {
	e.GET("/debug/*", echo.WrapHandler(http.DefaultServeMux))
}

func (r *_Router) Register(e *echo.Echo) {
	// Register your routes here
	r.registerRootRoutes(e)

	// 2.0
	auth.Routes.Register(e)
}

var Router = &_Router{}
