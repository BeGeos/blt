package auth

import (
	"go-hex/internal/apps/apperrors"
	"go-hex/internal/apps/auth"
	appjwt "go-hex/internal/infrastructure/jwt"
	"net/http"
)

var (
	ErrInvalidCredentials = apperrors.New("invalid_credentials", "invalid credentials", apperrors.WithHttpCode(http.StatusUnauthorized))
	ErrPersonNotFound     = apperrors.New("person_not_found", "person not found", apperrors.WithHttpCode(http.StatusNotFound))
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
	// TODO: set proper user validation logic
	if email == "admin@mail.com" && password == "password" {
		accessToken, err := s.JwtService.NewAccessToken(69)
		if err != nil {
			return auth.Tokens{}, appjwt.ErrFailedToken
		}

		refreshToken, err := s.JwtService.NewRefreshToken(69, 2)
		if err != nil {
			return auth.Tokens{}, appjwt.ErrFailedToken
		}

		return auth.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}, nil
	}
	return auth.Tokens{}, ErrInvalidCredentials
}

func (s *Service) Refresh(userID, version int) (auth.AccessToken, apperrors.AppError) {
	// check version with user version
	if version != 2 {
		return auth.AccessToken{}, appjwt.ErrInvalidTokenVersion
	}

	token, err := s.JwtService.NewAccessToken(userID)
	if err != nil {
		return auth.AccessToken{}, appjwt.ErrFailedToken
	}

	return auth.AccessToken{AccessToken: token}, nil
}
