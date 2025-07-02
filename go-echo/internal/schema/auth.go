package schema

type AuthResponse struct {
	Status string `json:"status"`
	Token  string `json:"token,omitempty"`
}

type LoginResponse struct {
	Status       string `json:"status"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
