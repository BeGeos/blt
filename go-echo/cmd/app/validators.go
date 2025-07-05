package app

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type _Validators struct {
	Request *RequestValidator
}

type RequestValidator struct {
	Validator *validator.Validate
}

func (rv *RequestValidator) Validate(i any) error {
	if err := rv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

var Validators = &_Validators{
	Request: &RequestValidator{Validator: validator.New()},
}
