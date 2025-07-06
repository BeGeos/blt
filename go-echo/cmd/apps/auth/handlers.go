package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/BeGeos/go-echo/internal/schema"
)

type Handler interface {
	Login(c echo.Context) error
	Logout(c echo.Context) error
	Me(c echo.Context) error
	Register(c echo.Context) error
	ResetPasswordRequest(c echo.Context) error
	ResetPassword(c echo.Context) error
	Refresh(c echo.Context) error
}
type handler struct {
	services *Services
	response *Response
}

func (h *handler) Login(c echo.Context) error {
	// TODO: fetch person by email and check password

	person := &schema.Person{}
	user := person.New()

	tokens, err := h.services.Token.GetValidTokens(user.UserID, user.Version)
	if err != nil {
		return err
	}

	return c.JSON(
		http.StatusOK,
		h.response.Tokens(tokens.AccessToken, tokens.RefreshToken),
	)
}

func (h *handler) Logout(c echo.Context) error {
	return c.JSON(http.StatusNoContent, h.response.Empty())
}

func (h *handler) Register(c echo.Context) error {
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

	tokens, err := h.services.Token.GetValidTokens(user.UserID, user.Version)
	if err != nil {
		return err
	}

	return c.JSON(
		http.StatusOK,
		h.response.Tokens(tokens.AccessToken, tokens.RefreshToken),
	)
}

func (h *handler) Me(c echo.Context) error {
	person := &schema.Person{}
	user := person.New()

	return c.JSON(http.StatusOK, user)
}

func (h *handler) Refresh(c echo.Context) error {
	person, _ := c.Get("person").(schema.PersonSchema) // add casting when person model is defined

	token, err := h.services.Token.GetNewAccessToken(person.UserID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, h.response.Token(token))
}

func (h *handler) ResetPasswordRequest(c echo.Context) error {
	var req ResetPasswordRequestDto
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	// generate reset password code and send email
	// either person exists or not, we return 204 No Content
	return c.JSON(http.StatusNoContent, h.response.Empty())
}

func (h *handler) ResetPassword(c echo.Context) error {
	var req ResetPasswordDto
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	// look for token and fetch person
	return c.JSON(http.StatusNoContent, h.response.Empty())
}

func NewHandler(s *Services) Handler {
	return &handler{
		services: s,
		response: &Response{},
	}
}
