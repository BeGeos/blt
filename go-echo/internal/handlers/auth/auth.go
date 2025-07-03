package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/BeGeos/go-echo/internal/dto"
	"github.com/BeGeos/go-echo/internal/schema"
	auth_services "github.com/BeGeos/go-echo/internal/services/auth"
)

type AuthHandler struct {
	Auth *auth_services.AuthService
}

func (h *AuthHandler) Login(c echo.Context) error {
	tokens, err := h.Auth.GetValidTokens()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &schema.TokensResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}

func (h *AuthHandler) Logout(c echo.Context) error {
	// increment user token version or invalidate the token
	return c.JSON(http.StatusNoContent, map[string]string{})
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req dto.RegisterRequestDto
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	// TODO: add logic to register the user, e.g., save to database

	tokens, err := h.Auth.GetValidTokens()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &schema.TokensResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}

func (h *AuthHandler) Me(c echo.Context) error {
	person := &schema.SelfResponse{
		UserID:    1,
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	return c.JSON(http.StatusOK, person)
}

func (h *AuthHandler) Refresh(c echo.Context) error {
	person, _ := c.Get("person").(uint) // add casting when person model is defined

	token, err := h.Auth.RefreshToken(person)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &schema.TokenResponse{
		AccessToken: token,
	})
}

func (h *AuthHandler) ResetPasswordRequest(c echo.Context) error {
	var req dto.ResetPasswordRequestDto
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	// generate reset password code and send email
	// either person exists or not, we return 204 No Content
	return c.JSON(http.StatusNoContent, map[string]string{})
}

func (h *AuthHandler) ResetPassword(c echo.Context) error {
	var req dto.ResetPasswordDto
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	// look for token and fetch person
	return c.JSON(http.StatusNoContent, map[string]string{})
}
