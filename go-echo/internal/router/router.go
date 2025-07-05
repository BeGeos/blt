package router

import (
	"github.com/labstack/echo/v4"

	api "github.com/BeGeos/go-echo/internal/router/api"

	"github.com/BeGeos/go-echo/cmd/apps/auth"
)

type _Router struct{}

func (r *_Router) Register(e *echo.Echo) {
	// Register your routes here

	RegisterRootRoutes(e)

	api.RegisterApiRoutes(e)

	// 2.0
	auth.Routes.Register(e)
}

var Router = &_Router{}
