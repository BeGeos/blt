package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

	"github.com/BeGeos/go-echo/internal"
	"github.com/BeGeos/go-echo/internal/apperrors"
)

func main() {
	e := echo.New()

	// run setup based on environment
	internal.Setup(e)

	e.GET("/", func(c echo.Context) error {
		luyckyNumber := os.Getenv("LUCKY_NUMBER")
		return c.String(http.StatusOK, fmt.Sprintf("Hello, world, the lucky number is %s ", luyckyNumber))
	})

	e.HTTPErrorHandler = apperrors.ErrorHandler // custom error handler

	e.Logger.Fatal(e.Start(":1323"))
}
