package jwt

import (
	"go-hex/internal/config"
	"log"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func successHandler(c echo.Context) {
	log.Println("JWT token successfully verified")
}

func newClaimsFunc(c echo.Context) jwt.Claims {
	return new(JwtClaims)
}

func getJwtConfig(args GetConfigArgs) echojwt.Config {
	cfg, _ := config.Load()

	return echojwt.Config{
		Skipper: func(c echo.Context) bool {
			auth := c.Request().Header.Get(echo.HeaderAuthorization)
			return auth == "" && !args.authenticationRequired
		},
		SuccessHandler: successHandler,
		ContextKey:     cfg.Authentication.ContextKey,
		SigningKey:     []byte(cfg.Authentication.SigningKey),
		NewClaimsFunc:  newClaimsFunc,
	}
}
