package middlewares

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	auth_services "github.com/BeGeos/go-echo/internal/services/auth"
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
		type RefreshTokenRequest struct {
			Token string `json:"token" validate:"required"`
		}
		var req RefreshTokenRequest
		if err := c.Bind(&req); err != nil {
			return echo.ErrBadRequest
		}

		token, err := jwt.Parse(req.Token, func(*jwt.Token) (interface{}, error) {
			return settings.JwtSecretKey, nil
		})
		if err != nil || !token.Valid {
			return echo.ErrUnauthorized
		}

		claims, ok := token.Claims.(auth_services.RefreshJwtClaims)
		if !ok {
			return echo.ErrUnauthorized
		}

		userID := claims.UserID   // use this to fetch person from database
		version := claims.Version // use this to compare person version

		if version != 1 { // replace with actual version check logic
			return echo.ErrUnauthorized
		}

		c.Set("person", userID) // Set the person in the context
		return next(c)
	}
}

func Authenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		person := c.Get("person")

		if person == nil {
			return echo.ErrUnauthorized
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
