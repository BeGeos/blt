package apperrors

import "net/http"

type AppError struct {
	Code    string
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func New(code, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func (e *AppError) ToHttp() int {
	switch e.Code {
	case "person_not_found":
		return http.StatusNotFound
	case "invalid_credentials":
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
