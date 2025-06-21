package logger

import (
	"context"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func LogValuesFunc(c echo.Context, v middleware.RequestLoggerValues) error {
	if v.Error == nil {
		logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
			slog.String("uri", v.URI),
			slog.String("method", v.Method),
			slog.Int("status", v.Status),
		)
	} else {
		logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
			slog.String("uri", v.URI),
			slog.String("method", v.Method),
			slog.Int("status", v.Status),
			slog.String("err", v.Error.Error()),
		)
	}
	return nil
}
