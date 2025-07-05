package auth

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"

	"github.com/BeGeos/go-echo/internal/schema"
)

type _Middlewares struct{}

func (m *_Middlewares) Authenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		person, ok := c.Get("person").(schema.PersonSchema)

		if !ok {
			return echo.ErrUnauthorized
		}

		if hub := sentryecho.GetHubFromContext(c); hub != nil {
			hub.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetUser(sentry.User{
					ID:        strconv.Itoa(int(person.UserID)),
					Email:     person.Email,
					IPAddress: c.RealIP(),
				})
			})
		}

		return next(c)
	}
}

func (m *_Middlewares) NotAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		person := c.Get("person")

		if person != nil {
			return echo.ErrForbidden
		}
		return next(c)
	}
}

var Middlewares = &_Middlewares{}
