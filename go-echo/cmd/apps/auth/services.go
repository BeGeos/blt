package auth

import (
	"errors"

	"github.com/BeGeos/go-echo/cmd/apps/jwt"

	appjwt "github.com/BeGeos/go-echo/cmd/apps/jwt"
)

type (
	Services struct {
		JwtServices *appjwt.Services
	}
	_AuthTokenService struct {
		JwtServices *appjwt.Services
	}
	_PasswordService struct{}
)

func (s *Services) Token() *_AuthTokenService {
	return &_AuthTokenService{
		JwtServices: s.JwtServices,
	}
}

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
		token, err := s.JwtServices.Token().New(claims, jwt.NewArgs{Kind: "access"})
		ch <- ResultCh{token: token, err: err}
	}(accessCh)

	go func(ch chan ResultCh) {
		claims := jwt.RefreshTokenClaims{
			UserID:  userID,
			Version: version, // replace with actual version check logic
		}
		token, err := s.JwtServices.Token().New(claims, jwt.NewArgs{Kind: "refresh"})
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
	claims := jwt.AccessTokenClaims{
		UserID: userID,
	}

	token, err := s.JwtServices.Token().New(claims, jwt.NewArgs{Kind: "access"})
	if err != nil {
		return "", errors.New("failed to create token")
	}

	return token, nil
}

type ServicesArgs struct {
	JwtServices *appjwt.Services
}

func NewServices(args ServicesArgs) *Services {
	return &Services{
		JwtServices: args.JwtServices,
	}
}
