package router

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/BeGeos/go-echo/internal/config"
	handlers "github.com/BeGeos/go-echo/internal/handlers/auth"
	"github.com/BeGeos/go-echo/internal/middlewares"
	services "github.com/BeGeos/go-echo/internal/services/auth"
)

func RegisterAuthRoutes(e *echo.Echo) {
	g := e.Group("/auth") // limited to 5 requests per minute

	authHandler := &handlers.AuthHandler{
		Auth: &services.AuthService{},
	}

	// requires authentication
	authenticated := g.Group("", echojwt.WithConfig(config.JwtAuthenticatedConfig), middlewares.Authenticated)

	authenticated.GET("/me", authHandler.Me).Name = "auth:me"
	authenticated.POST("/logout", authHandler.Logout).Name = "auth:logout"

	// must not be authenticated
	notAuthenticated := g.Group(
		"",
		middleware.RateLimiterWithConfig(config.GetAuthenticationRateLimiterConfig()),
		echojwt.WithConfig(config.JwtMaybeAuthenticatedConfig),
		middlewares.NotAuthenticated,
	)

	notAuthenticated.POST("/login", authHandler.Login).Name = "auth:login"
	notAuthenticated.POST("/register", authHandler.Register).Name = "auth:register"
	notAuthenticated.POST("/reset-password", authHandler.ResetPasswordRequest).Name = "auth:reset-password-request"
	notAuthenticated.POST("/reset-password/:token", authHandler.ResetPassword).Name = "auth:reset-password"

	maybeAuthenticated := g.Group(
		"",
		middleware.RateLimiterWithConfig(config.GetAuthenticationRateLimiterConfig()),
		middlewares.ValidateRefreshToken,
	) // one-handler group for documentation purposes
	maybeAuthenticated.POST("/refresh", authHandler.Refresh).Name = "auth:refresh"
}
