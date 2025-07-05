package auth

import (
	"errors"

	"github.com/BeGeos/go-echo/cmd/apps/jwt"
)

type _Services struct{}

func (s *_Services) Token() *_AuthTokenService {
	return &_AuthTokenService{}
}

type _AuthTokenService struct{}

// TODO: pass args for the claims
func (s *_AuthTokenService) GetValidTokens(userID, version uint) (Tokens, error) {
	type ResultCh struct {
		token string
		err   error
	}

	accessCh := make(chan ResultCh)
	refreshCh := make(chan ResultCh)

	go func(ch chan ResultCh) {
		claims := jwt.AccessTokenClaims{
			UserID: userID,
		}
		token, err := jwt.Services.NewAccessToken(claims)
		ch <- ResultCh{token: token, err: err}
	}(accessCh)

	go func(ch chan ResultCh) {
		claims := jwt.RefreshTokenClaims{
			UserID:  userID,
			Version: version, // replace with actual version check logic
		}
		token, err := jwt.Services.NewRefreshToken(claims)
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

func (s *_AuthTokenService) RefreshToken(userID uint) (string, error) {
	claims := jwt.AccessTokenClaims{
		UserID: userID,
	}

	token, err := jwt.Services.NewAccessToken(claims)
	if err != nil {
		return "", errors.New("failed to create token")
	}

	return token, nil
}

var Services = &_Services{}
