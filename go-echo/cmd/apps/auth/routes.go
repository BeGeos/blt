package auth

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	appjwt "github.com/BeGeos/go-echo/cmd/apps/jwt"
	"github.com/BeGeos/go-echo/internal/config"
)

func RegisterAuthRoutes(e *echo.Echo) {
	handlers := NewHandlers(NewServices(ServicesArgs{
		JwtServices: appjwt.NewServices(),
	}))
	g := e.Group("/auth") // limited to 5 requests per minute

	// requires authentication
	authenticated := g.Group("", echojwt.WithConfig(Config.JwtAuthenticatedConfig), Authenticated)

	authenticated.GET("/me", handlers.Me).Name = "auth:me"
	authenticated.POST("/logout", handlers.Logout).Name = "auth:logout"

	// must not be authenticated
	notAuthenticated := g.Group(
		"",
		middleware.RateLimiterWithConfig(config.GetAuthenticationRateLimiterConfig()),
		echojwt.WithConfig(Config.JwtMaybeAuthenticatedConfig),
		NotAuthenticated,
	)

	notAuthenticated.POST("/login", handlers.Login).Name = "auth:login"
	notAuthenticated.POST("/register", handlers.Register).Name = "auth:register"
	notAuthenticated.POST("/reset-password", handlers.ResetPasswordRequest).Name = "auth:reset-password-request"
	notAuthenticated.POST("/reset-password/:token", handlers.ResetPassword).Name = "auth:reset-password"

	maybeAuthenticated := g.Group(
		"",
		middleware.RateLimiterWithConfig(config.GetAuthenticationRateLimiterConfig()),
		appjwt.ValidateRefreshToken,
	) // one-handler group for documentation purposes
	maybeAuthenticated.POST("/refresh", handlers.Refresh).Name = "auth:refresh"
}
