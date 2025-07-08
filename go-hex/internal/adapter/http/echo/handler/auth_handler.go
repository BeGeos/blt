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
	type Request struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
	}
	var req Request

	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	tokens, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		return echo.NewHTTPError(err.ToHttp(), err.Error())
	}
	return c.JSON(http.StatusOK, tokens)
}

func (h *AuthHandler) Logout(c echo.Context) error {
	return c.JSON(http.StatusOK, "Logout endpoint not implemented yet")
}

func (h *AuthHandler) Register(c echo.Context) error {
	return c.JSON(http.StatusOK, "Register endpoint not implemented yet")
}

func (h *AuthHandler) Refresh(c echo.Context) error {
	return c.JSON(http.StatusOK, "Register endpoint not implemented yet")
}

func NewAuthHandler(authService *auth.Service) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}
