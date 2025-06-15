package internal

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

func loadEnv(env string) error {
	err := godotenv.Load(env)
	if err != nil {
		return fmt.Errorf("error loading %s file: %w", env, err)
	}
	return nil
}

func Setup(env string, e *echo.Echo) error {
	fmt.Println("Setting up...")

	switch env {
	case "dev":
		setupDev(e)
	case "prod":
		setupProd()
	case "staging":
		setupStaging()
	}

	return errors.New("APP_ENV is not set or invalid")
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
