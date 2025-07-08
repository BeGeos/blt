package auth

import (
	"go-hex/internal/apps/apperrors"
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
