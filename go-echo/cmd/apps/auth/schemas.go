package auth

type token struct {
	AccessToken string `json:"access_token,omitempty"`
}

type tokens struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// Init
type (
	_Schemas  struct{}
	_Response struct{}
)

// Centralise response hub
func (s *_Schemas) Response() *_Response {
	return &_Response{}
}

// Return an acess token response
func (r *_Response) Token(accessToken string) token {
	return token{
		AccessToken: accessToken,
	}
}

// Return an access and refresh token response
func (r *_Response) Tokens(accessToken, refreshToken string) tokens {
	return tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func (r *_Response) Empty() map[string]any {
	return map[string]any{}
}

var Schemas = &_Schemas{}
