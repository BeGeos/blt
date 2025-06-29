package middlewares

import (
	"github.com/labstack/echo/v4"
)

func ValidateTokenVersion(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// I assume person and version are set in the context by the JWT middleware
		// person := c.Get("person")
		version, _ := c.Get("version").(uint)

		// include person.version when db
		if version != 1 {
			return echo.ErrUnauthorized
		}

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
