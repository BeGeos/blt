package core

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func RootHandler(c echo.Context) error {
	luyckyNumber := os.Getenv("LUCKY_NUMBER")
	return c.String(http.StatusOK, fmt.Sprintf("Hello, world, the lucky number is %s ", luyckyNumber))
}

func HealthCheckHandler(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
