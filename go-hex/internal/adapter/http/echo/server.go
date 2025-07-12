package echo

import (
	"go-hex/internal/adapter"
	"go-hex/internal/config"
	"go-hex/pkg/sentry"
	"log"
	"net/http"
	"time"

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

	// Activate Sentry for error tracking
	tracing := &sentry.Client{Options: sentry.GetSentryClientOptions(cfg.Env)}
	if err := tracing.Init(); err != nil {
		log.Fatalf("sentry client init error: %v", err)
	}

	// Init Echo
	e := echo.New()

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

	// Open handlers
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ready ✅")
	}).Name = "root:index"
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	}).Name = "root:ping"

	e.HTTPErrorHandler = ErrorHandler
	e.Validator = NewValidator()

	// Add not debug setup
	if cfg.Env != config.EnvDevelopment {
		sentry.AddSentryMiddleware(e) // add sentry mdlw
	}

	RegisterRoutes(e)

	return e.Start(cfg.Server.Port)
}

func NewServer() adapter.Adapter {
	return &EchoServer{}
}
