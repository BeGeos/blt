package handlers

import (
	"errors"
	"net/http"

	"github.com/BeGeos/go-echo/internal/schema"
	auth_services "github.com/BeGeos/go-echo/internal/services/auth"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	Jwt *auth_services.JwtService
}

func (h *AuthHandler) Login(c echo.Context) error {
	// implement your login logic here

	claims := auth_services.AccessTokenClaims{
		UserID: 1,
	}
	token, err := h.Jwt.NewAccessToken(claims)
	if err != nil {
		return errors.New("Failed to create token")
	}

	return c.JSON(http.StatusOK, &schema.AuthResponse{
		Status: "success",
		Token:  token,
	})
}

func (h *AuthHandler) Logout(c echo.Context) error {
	// implement your login logic here
	// increment user token version or invalidate the token
	return c.JSON(http.StatusOK, &schema.AuthResponse{
		Status: "success",
	})
}

func (h *AuthHandler) Register(c echo.Context) error {
	// implement your login logic here
	return c.JSON(http.StatusOK, &schema.AuthResponse{
		Status: "success",
		Token:  "example",
	})
}

func (h *AuthHandler) Me(c echo.Context) error {
	// implement your login logic here
	return c.JSON(http.StatusOK, &schema.AuthResponse{
		Status: "success",
		Token:  "example",
	})
}

func (h *AuthHandler) Refresh(c echo.Context) error {
	// implement your login logic here
	return c.JSON(http.StatusOK, &schema.AuthResponse{
		Status: "success",
		Token:  "example",
	})
}
