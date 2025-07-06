package logger

import (
	"context"
	"go-hex/internal/config"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getLoggerLevel(env string) slog.Level {
	if env == "dev" {
		return slog.LevelDebug
	}

	return slog.LevelInfo // Default to Info level for production and staging
}

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func GetLogger(name string) *slog.Logger {
	cfg, _ := config.Load()
	return slog.New(
		slog.NewJSONHandler(
			os.Stderr,
			&slog.HandlerOptions{Level: getLoggerLevel(cfg.Env)},
		),
	)
}

func LogValuesFunc(c echo.Context, v middleware.RequestLoggerValues) error {
	request := slog.Group("request", slog.String("uri", v.URI), slog.String("method", v.Method), slog.Int("status", v.Status))
	metadata := slog.Group(
		"metadata",
		slog.String("user-agent", v.UserAgent),
		slog.String("protocol", v.Protocol),
		slog.String("remote_ip", v.RemoteIP),
		slog.String("request_id", v.RequestID),
		slog.Duration("latency", v.Latency))

	if v.Error == nil {
		logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
			request,
			metadata,
		)
	} else {
		logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
			request,
			metadata,
			slog.String("err", v.Error.Error()),
		)
	}
	return nil
}
