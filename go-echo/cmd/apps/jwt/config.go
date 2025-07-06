package jwt

import (
	"log"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/BeGeos/go-echo/internal/schema"
	"github.com/BeGeos/go-echo/internal/settings"
)

type _Config struct {
	JwtAuthenticatedConfig      echojwt.Config
	JwtMaybeAuthenticatedConfig echojwt.Config
}

var contextKey = settings.JwtContextKey

type Args struct {
	authenticationRequired bool
}

func successHandler(c echo.Context) {
	person := &schema.Person{}
	user := person.New()

	token, ok := c.Get(contextKey).(*jwt.Token)
	if ok {
		claims, _ := token.Claims.(*JwtClaims)

		// fetch person from database and put into context
		// TODO: remove this print
		log.Printf("the user %d has been successfully verified", claims.UserID)
		c.Set("person", user)
	}
}

func newClaimsFunc(c echo.Context) jwt.Claims {
	return new(JwtClaims)
}

func getJwtConfig(args Args) echojwt.Config {
	return echojwt.Config{
		Skipper: func(c echo.Context) bool {
			auth := c.Request().Header.Get(echo.HeaderAuthorization)
			return auth == "" && !args.authenticationRequired
		},
		SuccessHandler: successHandler,
		ContextKey:     contextKey,
		SigningKey:     []byte(settings.JwtSecretKey),
		NewClaimsFunc:  newClaimsFunc,
	}
}

var (
	jwtAuthenticatedConfig      = getJwtConfig(Args{authenticationRequired: true})
	jwtMaybeAuthenticatedConfig = getJwtConfig(Args{authenticationRequired: false})
)

var Config = &_Config{
	JwtAuthenticatedConfig:      jwtAuthenticatedConfig,
	JwtMaybeAuthenticatedConfig: jwtMaybeAuthenticatedConfig,
}
