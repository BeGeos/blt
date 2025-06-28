package schema

type AuthResponse struct {
	Status string `json:"status"`
	Token  string `json:"token,omitempty"`
}
