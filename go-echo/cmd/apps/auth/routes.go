package auth

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/BeGeos/go-echo/cmd/apps/jwt"
	"github.com/BeGeos/go-echo/internal/config"
)

type _Routes struct{}

func (r *_Routes) Register(e *echo.Echo) {
	g := e.Group("/auth") // limited to 5 requests per minute

	// requires authentication
	authenticated := g.Group("", echojwt.WithConfig(Config.JwtAuthenticatedConfig), Middlewares.Authenticated)

	authenticated.GET("/me", Handlers.Me).Name = "auth:me"
	authenticated.POST("/logout", Handlers.Logout).Name = "auth:logout"

	// must not be authenticated
	notAuthenticated := g.Group(
		"",
		middleware.RateLimiterWithConfig(config.GetAuthenticationRateLimiterConfig()),
		echojwt.WithConfig(Config.JwtMaybeAuthenticatedConfig),
		Middlewares.NotAuthenticated,
	)

	notAuthenticated.POST("/login", Handlers.Login).Name = "auth:login"
	notAuthenticated.POST("/register", Handlers.Register).Name = "auth:register"
	notAuthenticated.POST("/reset-password", Handlers.ResetPasswordRequest).Name = "auth:reset-password-request"
	notAuthenticated.POST("/reset-password/:token", Handlers.ResetPassword).Name = "auth:reset-password"

	maybeAuthenticated := g.Group(
		"",
		middleware.RateLimiterWithConfig(config.GetAuthenticationRateLimiterConfig()),
		jwt.Middlewares.ValidateRefreshToken,
	) // one-handler group for documentation purposes
	maybeAuthenticated.POST("/refresh", Handlers.Refresh).Name = "auth:refresh"
}

var Routes = &_Routes{}
