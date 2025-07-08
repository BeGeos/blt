package auth

import (
	"go-hex/internal/apps/apperrors"
	"go-hex/internal/apps/auth"
	appjwt "go-hex/internal/infrastructure/jwt"
)

var (
	ErrInvalidCredentials = apperrors.New("invalid_credentials", "invalid credentials")
	ErrPersonNotFound     = apperrors.New("person_not_found", "person not found")
)

type Service struct {
	JwtService *appjwt.JwtService
}

func NewAuthService(jwtService *appjwt.JwtService) *Service {
	return &Service{
		JwtService: jwtService,
	}
}

func (s *Service) Login(email, password string) (auth.Tokens, apperrors.AppError) {
	type ResultCh struct {
		token string
		err   error
	}
	accessCh := make(chan ResultCh)
	refreshCh := make(chan ResultCh)

	// TODO: set proper user validation logic
	if email == "admin@mail.com" && password == "password" {
		go func(ch chan ResultCh) {
			claims := appjwt.AccessTokenClaims{
				UserID: 69,
			}
			token, err := s.JwtService.NewAccessToken(claims)
			ch <- ResultCh{token: token, err: err}
		}(accessCh)

		go func(ch chan ResultCh) {
			claims := appjwt.RefreshTokenClaims{
				UserID:  69,
				Version: 1, // replace with actual version check logic
			}
			token, err := s.JwtService.NewRefreshToken(claims)
			ch <- ResultCh{token: token, err: err}
		}(refreshCh)

		accessResult := <-accessCh
		refreshResult := <-refreshCh

		if accessResult.err != nil || refreshResult.err != nil {
			return auth.Tokens{}, apperrors.New("generic_error", "failed to create token")
		}

		return auth.Tokens{
			AccessToken:  accessResult.token,
			RefreshToken: refreshResult.token,
		}, nil
	}
	return auth.Tokens{}, ErrInvalidCredentials
}
