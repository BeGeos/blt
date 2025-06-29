package config

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	auth_services "github.com/BeGeos/go-echo/internal/services/auth"
	"github.com/BeGeos/go-echo/internal/settings"
)

type Args struct {
	authenticationRequired bool
}

func successHandler(c echo.Context) {
}

func newClaimsFunc(c echo.Context) jwt.Claims {
	return new(auth_services.JwtClaims)
}

func getJwtConfig(args Args) echojwt.Config {
	return echojwt.Config{
		Skipper: func(c echo.Context) bool {
			auth := c.Request().Header.Get(echo.HeaderAuthorization)
			return auth == "" && !args.authenticationRequired
		},
		SuccessHandler: successHandler,
		ContextKey:     "tkn",
		SigningKey:     []byte(settings.JwtSecretKey),
		NewClaimsFunc:  newClaimsFunc,
	}
}

var (
	JwtAuthenticatedConfig      = getJwtConfig(Args{authenticationRequired: true})
	JwtMaybeAuthenticatedConfig = getJwtConfig(Args{authenticationRequired: false})
)
