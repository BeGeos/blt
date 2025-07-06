package auth

import (
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"

	appjwt "github.com/BeGeos/go-echo/cmd/apps/auth/jwt"
	"github.com/BeGeos/go-echo/internal/schema"
	"github.com/BeGeos/go-echo/internal/settings"
)

func Authenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		person, ok := c.Get("person").(schema.PersonSchema)

		if !ok {
			return echo.ErrUnauthorized
		}

		if hub := sentryecho.GetHubFromContext(c); hub != nil {
			hub.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetUser(sentry.User{
					ID:        strconv.Itoa(int(person.UserID)),
					IPAddress: c.RealIP(),
				})
			})
		}

		return next(c)
	}
}

func NotAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		person := c.Get("person")

		if person != nil {
			return echo.ErrForbidden
		}
		return next(c)
	}
}

// Validation happens in 4 steps:
//   - validate the request body => !ok return 401
//   - validate the token => !ok return 401
//   - validate the claims => !ok return 401
//   - validate the person version => !ok return 401
//
// If everything is ok, set the person in the context and call next handler
func ValidateRefreshToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req RefreshTokenRequestDto
		if err := c.Bind(&req); err != nil {
			return echo.ErrBadRequest
		}

		if err := c.Validate(&req); err != nil {
			return err
		}

		token, err := jwt.ParseWithClaims(
			req.Token,
			&appjwt.RefreshJwtClaims{},
			func(*jwt.Token) (interface{}, error) {
				return []byte(settings.JwtSecretKey), nil
			})
		if err != nil || !token.Valid {
			return echo.ErrUnauthorized
		}

		claims, ok := token.Claims.(*appjwt.RefreshJwtClaims)
		if !ok {
			return echo.ErrUnauthorized
		}

		version := claims.Version // use this to compare person version

		if version != 1 { // replace with actual version check logic
			return echo.ErrUnauthorized
		}

		person := &schema.Person{}
		user := person.New()

		c.Set("person", user) // Set the person in the context
		return next(c)
	}
}
