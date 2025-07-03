package schema

type SelfResponse struct {
	// tmp -> remove when person is on db
	UserID    uint   `json:"user_id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Age       int    `json:"age,omitempty"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token,omitempty"`
}

type TokensResponse struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
