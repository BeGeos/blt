package router

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/BeGeos/go-echo/internal/config"
	handlers "github.com/BeGeos/go-echo/internal/handlers/auth"
	"github.com/BeGeos/go-echo/internal/middlewares"
	services "github.com/BeGeos/go-echo/internal/services/auth"
)

func RegisterAuthRoutes(e *echo.Echo) {
	g := e.Group("/auth")

	authHandler := &handlers.AuthHandler{
		Auth: &services.AuthService{},
	}

	// requires authentication
	authenticated := g.Group("", echojwt.WithConfig(config.JwtAuthenticatedConfig), middlewares.Authenticated)

	authenticated.GET("/me", authHandler.Me).Name = "auth:me"
	authenticated.POST("/logout", authHandler.Logout).Name = "auth:logout"

	// must not be authenticated
	notAuthenticated := g.Group("", echojwt.WithConfig(config.JwtMaybeAuthenticatedConfig), middlewares.NotAuthenticated)

	notAuthenticated.POST("/login", authHandler.Login).Name = "auth:login"
	notAuthenticated.POST("/register", authHandler.Register).Name = "auth:register"

	maybeAuthenticated := g.Group("", middlewares.ValidateRefreshToken) // one-handler group for documentation purposes
	maybeAuthenticated.POST("/refresh", authHandler.Refresh).Name = "auth:refresh"
}
