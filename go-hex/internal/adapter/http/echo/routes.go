package echo

import (
	appecho "go-hex/internal/adapter/http/echo/handler"
	appmiddleware "go-hex/internal/adapter/http/echo/middleware"
	auth "go-hex/internal/apps/auth/app"
	appjwt "go-hex/internal/infrastructure/jwt"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func registerAuthRoutes(e *echo.Echo) {
	h := appecho.NewAuthHandler(
		auth.NewAuthService(
			appjwt.NewJwtService(),
		),
	)
	mcfg := appmiddleware.NewMiddlewareConfig()

	g := e.Group("/auth")

	authenticated := g.Group("", echojwt.WithConfig(
		appjwt.GetJwtConfig(appjwt.GetConfigArgs{AuthenticationRequired: true})),
		appmiddleware.Authenticated,
	)
	authenticated.POST("/logout", h.Logout).Name = "auth:logout"

	notAuthenticated := g.Group("",
		middleware.RateLimiterWithConfig(mcfg.AuthenticationRateLimiter()),
		echojwt.WithConfig(
			appjwt.GetJwtConfig(appjwt.GetConfigArgs{AuthenticationRequired: false})),
		appmiddleware.NotAuthenticated,
	)
	notAuthenticated.POST("/login", h.Login).Name = "auth:login"
	notAuthenticated.POST("/register", h.Register).Name = "auth:register"

	maybeAuthenticated := g.Group("",
		middleware.RateLimiterWithConfig(mcfg.AuthenticationRateLimiter()),
	) // group for docs only
	maybeAuthenticated.POST("/refresh", h.Refresh).Name = "auth:refresh"
}

func RegisterRoutes(e *echo.Echo) {
	registerAuthRoutes(e)
}
