package logger

import (
	"go-hex/internal/config"
	"log/slog"
	"os"
)

func getLogger(env string) *slog.Logger {
	if env == "dev" {
		return slog.New(
			slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}

	return slog.New(
		slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}),
	)
}

func GetLogger(name string) *slog.Logger {
	cfg, _ := config.Load()
	return getLogger(cfg.Env)
}

func GetRequestLogger() *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			nil,
		),
	)
}
