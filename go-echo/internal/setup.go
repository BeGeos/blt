package internal

import (
	"errors"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

func Setup(env string, e *echo.Echo) error {
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
	godotenv.Load(".env.local")

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
	godotenv.Load(".env")
}

func setupStaging() {
	log.Println("Setting up staging environment")
	godotenv.Load(".env.staging")
}
