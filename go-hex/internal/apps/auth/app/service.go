package auth

import (
	"go-hex/internal/apps/apperrors"
	"go-hex/internal/apps/auth"
	appjwt "go-hex/internal/infrastructure/jwt"
)

var (
	ErrInvalidCredentials  = apperrors.New("invalid_credentials", "invalid credentials")
	ErrPersonNotFound      = apperrors.New("person_not_found", "person not found")
	ErrInvalidTokenVersion = apperrors.New("invalid_token_version", "person not found")
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
			return auth.Tokens{}, apperrors.New("generic_error", "failed to create access token")
		}

		refreshToken, err := s.JwtService.NewRefreshToken(69, 2)
		if err != nil {
			return auth.Tokens{}, apperrors.New("generic_error", "failed to create refresh token")
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
		return auth.AccessToken{}, apperrors.New("invalid_token_version", "token is invalid or expired") // keep message generic but internal code clear
	}

	token, err := s.JwtService.NewAccessToken(userID)
	if err != nil {
		return auth.AccessToken{}, apperrors.New("generic_error", "failed to create access token")
	}

	return auth.AccessToken{AccessToken: token}, nil
}
