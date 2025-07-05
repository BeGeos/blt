package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"github.com/BeGeos/go-echo/internal/schema"
	"github.com/BeGeos/go-echo/internal/settings"
)

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
			&RefreshJwtClaims{},
			func(*jwt.Token) (interface{}, error) {
				return []byte(settings.JwtSecretKey), nil
			})
		if err != nil || !token.Valid {
			return echo.ErrUnauthorized
		}

		claims, ok := token.Claims.(*RefreshJwtClaims)
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
