package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/BeGeos/go-echo/internal/schema"
)

type _Handlers struct{}

func (h *_Handlers) Login(c echo.Context) error {
	// TODO: fetch person by email and check password

	person := &schema.Person{}
	user := person.New()

	tokens, err := Services.Token().GetValidTokens(user.UserID, user.Version)
	if err != nil {
		return err
	}

	return c.JSON(
		http.StatusOK,
		Schemas.Response().Tokens(tokens.AccessToken, tokens.RefreshToken),
	)
}

func (h *_Handlers) Logout(c echo.Context) error {
	return c.JSON(http.StatusNoContent, Schemas.Response().Empty())
}

func (h *_Handlers) Register(c echo.Context) error {
	var req RegisterRequestDto
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	// TODO: add logic to register the user, e.g., save to database
	person := &schema.Person{}
	user := person.New()

	tokens, err := Services.Token().GetValidTokens(user.UserID, user.Version)
	if err != nil {
		return err
	}

	return c.JSON(
		http.StatusOK,
		Schemas.Response().Tokens(tokens.AccessToken, tokens.RefreshToken),
	)
}

func (h *_Handlers) Me(c echo.Context) error {
	person := &schema.Person{}
	user := person.New()

	return c.JSON(http.StatusOK, user)
}

func (h *_Handlers) Refresh(c echo.Context) error {
	person, _ := c.Get("person").(schema.PersonSchema) // add casting when person model is defined

	token, err := Services.Token().RefreshToken(person.UserID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, Schemas.Response().Token(token))
}

func (h *_Handlers) ResetPasswordRequest(c echo.Context) error {
	var req ResetPasswordRequestDto
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	// generate reset password code and send email
	// either person exists or not, we return 204 No Content
	return c.JSON(http.StatusNoContent, Schemas.Response().Empty())
}

func (h *_Handlers) ResetPassword(c echo.Context) error {
	var req ResetPasswordDto
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	// look for token and fetch person
	return c.JSON(http.StatusNoContent, Schemas.Response().Empty())
}

var Handlers = &_Handlers{}
