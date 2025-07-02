package app

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/time/rate"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/BeGeos/go-echo/internal/config"
	"github.com/BeGeos/go-echo/internal/handlers"
	"github.com/BeGeos/go-echo/internal/router"
	"github.com/BeGeos/go-echo/internal/settings"
	"github.com/BeGeos/go-echo/internal/validators"
)

func checkEnvVariables() error {
	// these are the variables that must be set in the environment
	return nil
}

func Setup() *echo.Echo {
	e := echo.New()

	fmt.Println("Setting up...")
	err := settings.LoadEnv()
	if err != nil {
		log.Fatalf("❌ Failed to load environment variables: %v\n", err)
	}
	fmt.Print("1. Environment variables loaded successfully ✅\n")

	if err := checkEnvVariables(); err != nil {
		log.Fatalf("❌ Failed to check env variables: %v\n", err)
	}
	fmt.Print("2. Environment variables checked successfully ✅\n")

	e.Use(middleware.RequestID())
	e.Use(middleware.RequestLoggerWithConfig(config.LoggerConfigs))
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("5M"))
	e.Use(middleware.CORSWithConfig(config.CorsConfig))
	e.Use(middleware.Gzip())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(30)))) // 30 requests per second
	e.Use(middleware.Secure())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second, // 30 seconds timeout
	}))

	e.HTTPErrorHandler = handlers.ErrorHandler // custom error handler
	e.Validator = &validators.RequestValidator{Validator: validator.New()}

	router.RegisterRoutes(e) // register routes

	return e
}
