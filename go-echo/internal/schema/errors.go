package schema

type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`  // optional
	Stack   string `json:"stack,omitempty"` // optional
}
