package echo

import (
	appecho "go-hex/internal/adapter/http/echo/handler"
	auth "go-hex/internal/apps/auth/app"
	appjwt "go-hex/internal/infrastructure/jwt"

	"github.com/labstack/echo/v4"
)

func registerAuthRoutes(e *echo.Echo) {
	h := appecho.NewAuthHandler(
		auth.NewAuthService(
			appjwt.NewJwtService(),
		),
	)

	g := e.Group("/auth")

	g.POST("/login", h.Login)
	g.POST("/logout", h.Logout)
	g.POST("/register", h.Register)
}

type Services struct {
	authService *auth.Service
}

func RegisterRoutes(e *echo.Echo) {
	registerAuthRoutes(e)
}
