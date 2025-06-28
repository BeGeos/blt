package v1_handlers

import (
	"net/http"

	"github.com/BeGeos/go-echo/internal/schema"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h *UserHandler) List(c echo.Context) error {
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	users := []User{
		{Name: "Alice", Email: "alice@example.com"},
		{Name: "Bob", Email: "bob@example.com"},
		{Name: "Charlie", Email: "charlie@example.com"},
		{Name: "Diana", Email: "diana@example.com"},
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) Get(c echo.Context) error {
	userID := c.Param("id")
	return c.JSON(http.StatusOK, schema.BaseResponse{Message: "this is user " + userID})
}
