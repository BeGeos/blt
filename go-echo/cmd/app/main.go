package app

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"golang.org/x/time/rate"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/BeGeos/go-echo/internal/config"
	"github.com/BeGeos/go-echo/internal/handlers"
	"github.com/BeGeos/go-echo/internal/settings"
	"github.com/BeGeos/go-echo/pkg/sentry"
)

func checkEnvVariables() error {
	// these are the variables that must be set in the environment
	return nil
}

func addPprofRoutes(e *echo.Echo) {
	e.GET("/debug/*", echo.WrapHandler(http.DefaultServeMux))
}

type Stepper struct {
	start *int
}

func (s *Stepper) increment() {
	*s.start++
}

func (s *Stepper) decrement() {
	*s.start--
}

func Setup() *echo.Echo {
	env, step := settings.Env, 1
	stepper := &Stepper{start: &step}

	fmt.Println("Setting up...")

	// Activate Sentry for error tracking
	tracing := &sentry.Client{Options: sentry.GetSentryClientOptions(env)}
	if err := tracing.Init(); err != nil {
		log.Fatalf("sentry client init error: %v", err)
	}
	fmt.Printf("%d. Sentry Client init successfully ✅\n", step)
	stepper.increment()

	e := echo.New()

	err := settings.LoadEnv()
	if err != nil {
		log.Fatalf("❌ Failed to load environment variables: %v\n", err)
	}
	fmt.Printf("%d. Environment variables loaded successfully ✅\n", step)
	stepper.increment()

	if err := checkEnvVariables(); err != nil {
		log.Fatalf("❌ Failed to check env variables: %v\n", err)
	}
	fmt.Printf("%d. Environment variables checked successfully ✅\n", step)
	stepper.increment()

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
	e.Validator = Validators.Request           // custom request validator

	// Add not debug setup
	if env != settings.EnvDevelopment {
		sentry.AddSentryMiddleware(e) // add sentry
		fmt.Printf("%d. Added Sentry middleware ✅\n", step)
		stepper.increment()
	}

	Router.Register(e) // register routes

	// Add dev setup
	if env == settings.EnvDevelopment {
		addPprofRoutes(e) // add pprof routes for debugging
		fmt.Printf("%d. Pprof routes added for debugging ✅\n", step)
		stepper.increment()

	}

	return e
}
