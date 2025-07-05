package services

import (
	"errors"
)

type AuthService struct{}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

// TODO: pass args for the claims
func (s *AuthService) GetValidTokens(userID, version uint) (Tokens, error) {
	jwt := &JwtService{}

	type ResultCh struct {
		token string
		err   error
	}

	accessCh := make(chan ResultCh)
	refreshCh := make(chan ResultCh)

	go func(ch chan ResultCh) {
		claims := AccessTokenClaims{
			UserID: userID,
		}
		token, err := jwt.NewAccessToken(claims)
		ch <- ResultCh{token: token, err: err}
	}(accessCh)

	go func(ch chan ResultCh) {
		claims := RefreshTokenClaims{
			UserID:  userID,
			Version: version, // replace with actual version check logic
		}
		token, err := jwt.NewRefreshToken(claims)
		ch <- ResultCh{token: token, err: err}
	}(refreshCh)

	accessResult := <-accessCh
	refreshResult := <-refreshCh

	if accessResult.err != nil || refreshResult.err != nil {
		return Tokens{}, errors.New("failed to create token")
	}

	return Tokens{
		AccessToken:  accessResult.token,
		RefreshToken: refreshResult.token,
	}, nil
}

func (s *AuthService) RefreshToken(userID uint) (string, error) {
	jwt := &JwtService{}

	claims := AccessTokenClaims{
		UserID: userID,
	}

	token, err := jwt.NewAccessToken(claims)
	if err != nil {
		return "", errors.New("failed to create token")
	}

	return token, nil
}
