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
	// TODO: fetch person by email and check password

	person := &schema.Person{}
	user := person.New()

	tokens, err := h.Auth.GetValidTokens(user.UserID, user.Version)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &schema.TokensResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}

func (h *AuthHandler) Logout(c echo.Context) error {
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
	person := &schema.Person{}
	user := person.New()

	tokens, err := h.Auth.GetValidTokens(user.UserID, user.Version)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &schema.TokensResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}

func (h *AuthHandler) Me(c echo.Context) error {
	person := &schema.Person{}
	user := person.New()

	panic("this one comes with request id")

	return c.JSON(http.StatusOK, user)
}

func (h *AuthHandler) Refresh(c echo.Context) error {
	person, _ := c.Get("person").(schema.PersonSchema) // add casting when person model is defined

	token, err := h.Auth.RefreshToken(person.UserID)
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
