package logger

import (
	"go-hex/internal/config"
	"log/slog"
	"os"
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

func GetRequestLogger() *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			nil,
		),
	)
}
