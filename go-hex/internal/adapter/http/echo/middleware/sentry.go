package echo

import (
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
)

var MiddlewareOptions = sentryecho.Options{
	Repanic: true,
}

func AddSentryMiddleware(e *echo.Echo) {
	e.Use(sentryecho.New(MiddlewareOptions))
}
