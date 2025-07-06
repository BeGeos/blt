package auth

type token struct {
	AccessToken string `json:"access_token,omitempty"`
}

type tokens struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type Response struct{}

// Return an acess token response
func (r *Response) Token(accessToken string) token {
	return token{
		AccessToken: accessToken,
	}
}

// Return an access and refresh token response
func (r *Response) Tokens(accessToken, refreshToken string) tokens {
	return tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func (r *Response) Empty() map[string]any {
	return map[string]any{}
}
