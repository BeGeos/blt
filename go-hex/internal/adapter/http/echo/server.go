package echo

import (
	"go-hex/internal/config"
	"net/http"

	apphttp "go-hex/internal/adapter/http"

	"github.com/labstack/echo/v4"
)

type EchoServer struct{}

func (s *EchoServer) Start() error {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	// Init Echo
	e := echo.New()

	// Open handlers
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ready ✅")
	})
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	return e.Start(cfg.Server.Port)
}

func NewServer() apphttp.Server {
	return &EchoServer{}
}
