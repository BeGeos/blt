package v1_router

import (
	v1_handlers "github.com/BeGeos/go-echo/internal/handlers/api/v1"
	"github.com/labstack/echo/v4"
)

func RegisterV1Routes(e *echo.Group) {
	g := e.Group("/v1")

	// Root handler for v1
	g.GET("", v1_handlers.RootHandler)
}
