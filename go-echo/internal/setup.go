package internal

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/time/rate"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func loadEnv(env string) error {
	err := godotenv.Load(env)
	if err != nil {
		return fmt.Errorf("error loading %s file: %w", env, err)
	}
	return nil
}

func checkEnvVariabled() error {
	// these are the variables that must be set in the environment
	Vars := []string{}
	return nil
}

func setupDev(e *echo.Echo) {
	log.Println("Setting up development environment")
	err := loadEnv(".env.local")
	if err != nil {
		log.Println("❌ Failed to load env: %v\n", err)
		os.Exit(1)
	}

	e.Use(echoprometheus.NewMiddleware("go-echo")) // adds middleware to gather metrics
	go func() {
		metrics := echo.New()                                // this Echo will run on separate port 8081
		metrics.GET("/metrics", echoprometheus.NewHandler()) // adds route to serve gathered metrics
		if err := metrics.Start(":8081"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()
}

func setupProd() {
	log.Println("Setting up production environment")

	err := loadEnv(".env")
	if err != nil {
		log.Println("❌ Failed to load env: %v\n", err)
		os.Exit(1)
	}
}

func setupStaging() {
	log.Println("Setting up staging environment")

	err := loadEnv(".env.staging")
	if err != nil {
		log.Println("❌ Failed to load env: %v\n", err)
		os.Exit(1)
	}
}

func Setup(e *echo.Echo) error {
	Env := os.Getenv("APP_ENV")
	fmt.Println("Setting up...")

	switch Env {
	case "dev":
		setupDev(e)
	case "prod":
		setupProd()
	case "staging":
		setupStaging()
	default:
		return errors.New("APP_ENV is not set or invalid")
	}

	if err := checkEnvVariabled(); err != nil {
		log.Println("❌ Failed to check env variables: %v\n", err)
	}

	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("10M"))
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(30)))) // 30 requests per second
	e.Use(middleware.Secure())
	e.Use(middleware.Timeout())

	return nil
}
