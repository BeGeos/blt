package apperrors

type AppError interface {
	Error() string
	ToHttp() int
}

type appError struct {
	Code    string
	Message string
	Cause   error
	http    int
}

type Opts func(*appError)

func (e *appError) Error() string {
	return e.Message
}

func New(code, message string, opts ...Opts) AppError {
	ae := &appError{
		Code:    code,
		Message: message,
	}

	for _, opt := range opts {
		opt(ae)
	}

	return ae
}

func (e *appError) ToHttp() int {
	return e.http
}

func WithHttpCode(code int) Opts {
	return func(e *appError) {
		e.http = code
	}
}

func WithCause(err error) Opts {
	return func(e *appError) {
		e.Cause = err
	}
}
