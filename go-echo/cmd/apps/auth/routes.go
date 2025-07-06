package auth

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	appjwt "github.com/BeGeos/go-echo/cmd/apps/jwt"
	"github.com/BeGeos/go-echo/internal/config"
)

func RegisterAuthRoutes(e *echo.Echo) {
	handler := NewHandler(NewServices(ServicesArgs{
		JwtServices: appjwt.NewServices(),
	}))
	g := e.Group("/auth")

	// requires authentication
	authenticated := g.Group("", echojwt.WithConfig(appjwt.Config.JwtAuthenticatedConfig), Authenticated)

	authenticated.GET("/me", handler.Me).Name = "auth:me"
	authenticated.POST("/logout", handler.Logout).Name = "auth:logout"

	// must not be authenticated
	notAuthenticated := g.Group(
		"",
		middleware.RateLimiterWithConfig(config.GetAuthenticationRateLimiterConfig()), // limited to 5 requests per minute
		echojwt.WithConfig(appjwt.Config.JwtMaybeAuthenticatedConfig),
		NotAuthenticated,
	)

	notAuthenticated.POST("/login", handler.Login).Name = "auth:login"
	notAuthenticated.POST("/register", handler.Register).Name = "auth:register"
	notAuthenticated.POST("/reset-password", handler.ResetPasswordRequest).Name = "auth:reset-password-request"
	notAuthenticated.POST("/reset-password/:token", handler.ResetPassword).Name = "auth:reset-password"

	maybeAuthenticated := g.Group(
		"",
		middleware.RateLimiterWithConfig(config.GetAuthenticationRateLimiterConfig()),
		appjwt.ValidateRefreshToken,
	) // one-handler group for documentation purposes
	maybeAuthenticated.POST("/refresh", handler.Refresh).Name = "auth:refresh"
}
