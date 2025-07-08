package echo

import (
	auth "go-hex/internal/apps/auth/app"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService *auth.Service
}

func (h *AuthHandler) Login(c echo.Context) error {
	return c.JSON(http.StatusOK, "Login endpoint not implemented yet")
}

func (h *AuthHandler) Logout(c echo.Context) error {
	return c.JSON(http.StatusOK, "Logout endpoint not implemented yet")
}

func (h *AuthHandler) Register(c echo.Context) error {
	return c.JSON(http.StatusOK, "Register endpoint not implemented yet")
}

func NewAuthHandler(authService *auth.Service) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}
