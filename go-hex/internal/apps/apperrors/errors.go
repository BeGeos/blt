package apperrors

import "net/http"

type AppError interface {
	Error() string
	ToHttp() int
}

type appError struct {
	Code    string
	Message string
}

func (e *appError) Error() string {
	return e.Message
}

func New(code, message string) AppError {
	return &appError{
		Code:    code,
		Message: message,
	}
}

func (e *appError) ToHttp() int {
	switch e.Code {
	case "person_not_found":
		return http.StatusNotFound
	case "invalid_credentials":
		return http.StatusUnauthorized
	case "invalid_token_version":
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
