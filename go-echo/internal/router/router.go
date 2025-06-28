package router

import (
	"github.com/labstack/echo/v4"

	"github.com/BeGeos/go-echo/internal/router/api"
	"github.com/BeGeos/go-echo/internal/router/auth"
)

func RegisterRoutes(e *echo.Echo) {
	// Register your routes here

	RegisterRootRoutes(e)

	api.RegisterApiRoutes(e)
	auth.RegisterAuthRoutes(e)
}
