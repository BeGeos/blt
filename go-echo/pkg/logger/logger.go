package logger

import (
	"context"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/BeGeos/go-echo/internal/settings"
)

func getLoggerLevel() slog.Level {
	Env := settings.Env
	if Env == "dev" {
		return slog.LevelDebug
	}

	return slog.LevelInfo // Default to Info level for production and staging
}

var (
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	Logger = slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: getLoggerLevel()}))
)

func LogValuesFunc(c echo.Context, v middleware.RequestLoggerValues) error {
	request := slog.Group("request", slog.String("uri", v.URI), slog.String("method", v.Method), slog.Int("status", v.Status))

	if v.Error == nil {
		logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
			request,
		)
	} else {
		logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
			request,
			slog.String("err", v.Error.Error()),
		)
	}
	return nil
}
