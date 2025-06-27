package router

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// Register your routes here

	RegisterRootRoutes(e)
}
