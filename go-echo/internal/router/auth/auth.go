package auth

import (
	auth_handlers "github.com/BeGeos/go-echo/internal/handlers/auth"
	"github.com/labstack/echo/v4"
)

func RegisterAuthRoutes(e *echo.Echo) {
	g := e.Group("/auth")

	authHandler := &auth_handlers.AuthHandler{}

	g.GET("/me", authHandler.Me).Name = "auth:me"

	g.POST("/login", authHandler.Login).Name = "auth:login"
	g.POST("/logout", authHandler.Logout).Name = "auth:logout"
	g.POST("/register", authHandler.Register).Name = "auth:register"
	g.POST("/refresh", authHandler.Refresh).Name = "auth:refresh"
}
