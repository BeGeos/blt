package api

import (
	"github.com/labstack/echo/v4"

	v1_router "github.com/BeGeos/go-echo/internal/router/api/v1"
)

func RegisterApiRoutes(e *echo.Echo) {
	g := e.Group("/api")

	v1_router.RegisterV1Routes(g)
}
