package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"

	"github.com/BeGeos/go-echo/internal"
)

func main() {
	e := echo.New()
	Env := os.Getenv("APP_ENV")

	// run setup based on environment
	internal.Setup(Env, e)

	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("10M"))
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(30)))) // 30 requests per second
	e.Use(middleware.Secure())
	e.Use(middleware.Timeout())

	e.GET("/", func(c echo.Context) error {
		luyckyNumber := os.Getenv("LUCKY_NUMBER")
		return c.String(http.StatusOK, fmt.Sprintf("Hello, world, the lucky number is %s ", luyckyNumber))
	})

	e.Logger.Fatal(e.Start(":1323"))
}
