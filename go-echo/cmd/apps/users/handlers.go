package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type _Handlers struct{}

func (_h *_Handlers) Get(c echo.Context) error {
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	return c.JSON(http.StatusOK, User{
		Name: "Alice", Email: "alice@example.com",
	})
}

var Handlers = &_Handlers{}
