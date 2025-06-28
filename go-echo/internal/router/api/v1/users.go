package v1_router

import (
	v1_handlers "github.com/BeGeos/go-echo/internal/handlers/api/v1"
	"github.com/labstack/echo/v4"
)

func RegisterV1UsersRoutes(g *echo.Group) {
	userHandler := &v1_handlers.UserHandler{}

	g.GET("/users", userHandler.List).Name = "api-v1:list-users"
	g.GET("/users/:id", userHandler.Get).Name = "api-v1:get-user"
}
