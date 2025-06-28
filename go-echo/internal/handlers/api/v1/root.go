package v1_handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string `json:"message"`
}

func RootHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, Response{Message: "this is v1"})
}
