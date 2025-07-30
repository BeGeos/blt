package core

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"
)

const (
	ErrConflict     = "conflict"
	ErrInternal     = "internal"
	ErrInvalid      = "invalid"
	ErrNotFound     = "not_found"
	ErrUnauthorized = "unauthorized"
	ErrForbidden    = "forbidden"
)

type Error struct {
	// Machine-readable error code
	Code string

	// Human-readable error message
	Message string

	// Logical operation and nested errors
	Op  string
	Err error

	Severity slog.Level
}

func (e *Error) Error() string {
	var buf bytes.Buffer

	// Print the current operation in our stack, if any.
	if e.Op != "" {
		fmt.Fprintf(&buf, "%s: ", e.Op)
	}

	// If wrapping an error, print its Error() message.
	// Otherwise print the error code & message.
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			fmt.Fprintf(&buf, "<%s> ", e.Code)
		}
		buf.WriteString(e.Message)
	}
	return buf.String()
}

func ErrorCode(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Code != "" {
		return e.Code
	} else if ok && e.Err != nil {
		return ErrorCode(e.Err)
	}
	return ErrInternal
}

func ErrorMessage(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}
	return "An internal error has occurred"
}

var errToHttpStatusCode = map[string]int{
	ErrConflict:     http.StatusConflict,
	ErrInternal:     http.StatusInternalServerError,
	ErrInvalid:      http.StatusBadRequest,
	ErrNotFound:     http.StatusNotFound,
	ErrUnauthorized: http.StatusUnauthorized,
	ErrForbidden:    http.StatusForbidden,
}

func ErrorHttp(err error) int {
	if err == nil {
		return 0
	}

	code := ErrorCode(err)
	httpCode, ok := errToHttpStatusCode[code]
	if ok {
		return httpCode
	}

	return http.StatusInternalServerError
}
