package router

import (
	"github.com/labstack/echo/v4"
)

func RegisterV1Routes(g *echo.Group) {
	gg := g.Group("/v1")

	RegisterV1UsersRoutes(gg)
}
