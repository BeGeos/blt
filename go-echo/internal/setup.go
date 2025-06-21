package internal

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"golang.org/x/time/rate"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/BeGeos/go-echo/internal/configs"
	"github.com/BeGeos/go-echo/internal/utils"
)

func loadEnv() (string, error) {
	Env := utils.GetEnv()

	var envFile string
	switch Env {
	case "dev":
		envFile = ".env.local"
	case "prod":
		envFile = ".env"
	case "staging":
		envFile = ".env.staging"
	}

	err := godotenv.Load(envFile)
	return Env, err
}

func checkEnvVariables() error {
	// these are the variables that must be set in the environment
	return nil
}

func Setup(e *echo.Echo) error {
	fmt.Println("Setting up...")
	_, err := loadEnv()
	if err != nil {
		log.Fatalf("❌ Failed to load environment variables: %v\n", err)
	}
	fmt.Print("1. Environment variables loaded successfully ✅\n")

	if err := checkEnvVariables(); err != nil {
		log.Fatalf("❌ Failed to check env variables: %v\n", err)
	}
	fmt.Print("2. Environment variables checked successfully ✅\n")

	e.Use(middleware.RequestID())
	e.Use(middleware.RequestLoggerWithConfig(configs.LoggerConfigs))
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("10M"))
	e.Use(middleware.CORSWithConfig(configs.CorsConfig))
	e.Use(middleware.Gzip())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(30)))) // 30 requests per second
	e.Use(middleware.Secure())
	e.Use(middleware.Timeout())

	return nil
}
