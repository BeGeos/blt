package echo

import (
	"go-hex/internal/config"
	"net/http"
	"time"

	apphttp "go-hex/internal/adapter/http"
	appmiddleware "go-hex/internal/adapter/http/echo/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
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

	mcfg := appmiddleware.NewMiddlewareConfig()

	e.Use(middleware.RequestID())
	e.Use(middleware.RequestLoggerWithConfig(mcfg.Logger()))
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("5M"))
	e.Use(middleware.CORSWithConfig(mcfg.Cors()))
	e.Use(middleware.Gzip())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(30)))) // 30 requests per second
	e.Use(middleware.Secure())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second, // 30 seconds timeout
	}))

	return e.Start(cfg.Server.Port)
}

func NewServer() apphttp.Server {
	return &EchoServer{}
}
