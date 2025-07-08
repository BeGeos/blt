package echo

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`  // optional
	Stack   string `json:"stack,omitempty"` // optional
}

func ErrorHandler(err error, c echo.Context) {
	// Default values
	code := http.StatusInternalServerError
	msg := "Internal Server Error"

	// If it's an *echo.HTTPError, extract its code and message
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code

		switch m := he.Message.(type) {
		case string:
			msg = m
		case fmt.Stringer:
			msg = m.String()
		default:
			msg = fmt.Sprintf("%v", m)
		}
	}

	// Send a JSON response
	res := ErrorResponse{
		Error:   true,
		Message: msg,
		Code:    code,
	}

	if !c.Response().Committed {
		_ = c.JSON(code, res)
	}
}
