package router

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/BeGeos/go-echo/internal/config"
	handlers "github.com/BeGeos/go-echo/internal/handlers/auth"
	auth_services "github.com/BeGeos/go-echo/internal/services/auth"
)

func RegisterAuthRoutes(e *echo.Echo) {
	g := e.Group("/auth")

	authHandler := &handlers.AuthHandler{
		Jwt: &auth_services.JwtService{},
	}

	// requires authentication
	authenticated := g.Group("")
	authenticated.Use(echojwt.WithConfig(config.JwtAuthenticatedConfig))

	authenticated.GET("/me", authHandler.Me).Name = "auth:me"
	authenticated.POST("/refresh", authHandler.Refresh).Name = "auth:refresh"
	authenticated.POST("/logout", authHandler.Logout).Name = "auth:logout"

	// must not be authenticated
	notAuthenticated := g.Group("")
	notAuthenticated.Use(echojwt.WithConfig(config.JwtNotAuthenticatedConfig))

	notAuthenticated.POST("/login", authHandler.Login).Name = "auth:login"
	notAuthenticated.POST("/register", authHandler.Register).Name = "auth:register"
}
