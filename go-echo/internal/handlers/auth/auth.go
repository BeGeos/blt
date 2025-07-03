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
	type ResultCh struct {
		token string
		err   error
	}

	accessCh := make(chan ResultCh)
	refreshCh := make(chan ResultCh)

	go func(ch chan ResultCh) {
		claims := auth_services.AccessTokenClaims{
			UserID: 1,
		}
		token, err := h.Jwt.NewAccessToken(claims)
		ch <- ResultCh{token: token, err: err}
	}(accessCh)

	go func(ch chan ResultCh) {
		claims := auth_services.RefreshTokenClaims{
			UserID:  1,
			Version: 1, // replace with actual version check logic
		}
		token, err := h.Jwt.NewRefreshToken(claims)
		ch <- ResultCh{token: token, err: err}
	}(refreshCh)

	accessResult := <-accessCh
	refreshResult := <-refreshCh

	if accessResult.err != nil || refreshResult.err != nil {
		return errors.New("Failed to create token")
	}

	return c.JSON(http.StatusOK, &schema.LoginResponse{
		Status:       "success",
		AccessToken:  accessResult.token,
		RefreshToken: refreshResult.token,
	})
}

func (h *AuthHandler) Logout(c echo.Context) error {
	// implement your login logic here
	// increment user token version or invalidate the token
	return c.JSON(http.StatusNoContent, map[string]string{})
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
	person, _ := c.Get("person").(uint) // add casting when person model is defined

	const version = 1 // replace with actual version check logic when person
	claims := auth_services.AccessTokenClaims{
		UserID: person,
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
