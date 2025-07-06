package auth

import (
	"errors"

	appjwt "github.com/BeGeos/go-echo/cmd/apps/auth/jwt"
)

type Token interface {
	GetValidTokens(userID, version uint) (Tokens, error)
	GetNewAccessToken(userID uint) (string, error)
}
type (
	Services struct {
		// Injected services
		Token Token
	}
)

// Subservices - this are private and
// managed by the main service
type (
	_AuthTokenService struct {
		Jwt appjwt.Jwt
	}
	_PasswordService struct{}
)

func (s *_AuthTokenService) GetValidTokens(userID, version uint) (Tokens, error) {
	type ResultCh struct {
		token string
		err   error
	}

	accessCh := make(chan ResultCh)
	refreshCh := make(chan ResultCh)

	go func(ch chan ResultCh) {
		claims := appjwt.AccessTokenClaims{
			UserID: userID,
		}
		token, err := s.Jwt.New(claims, appjwt.NewJwtArgs{Kind: "access"})
		ch <- ResultCh{token: token, err: err}
	}(accessCh)

	go func(ch chan ResultCh) {
		claims := appjwt.RefreshTokenClaims{
			UserID:  userID,
			Version: version, // replace with actual version check logic
		}
		token, err := s.Jwt.New(claims, appjwt.NewJwtArgs{Kind: "refresh"})
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

func (s *_AuthTokenService) GetNewAccessToken(userID uint) (string, error) {
	claims := appjwt.AccessTokenClaims{
		UserID: userID,
	}

	token, err := s.Jwt.New(claims, appjwt.NewJwtArgs{Kind: "access"})
	if err != nil {
		return "", errors.New("failed to create token")
	}

	return token, nil
}

func newAuthTokenService() Token {
	return &_AuthTokenService{
		Jwt: appjwt.NewJwt(),
	}
}

func NewServices() *Services {
	return &Services{
		Token: newAuthTokenService(),
	}
}
