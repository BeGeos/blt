package echo

import (
	"context"
	"go-hex/pkg/logger"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var _logger = logger.GetRequestLogger()

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
		_logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
			request,
			metadata,
		)
	} else {
		_logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
			request,
			metadata,
			slog.String("err", v.Error.Error()),
		)
	}
	return nil
}
