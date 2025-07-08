package auth

import (
	"errors"

	appjwt "go-hex/internal/infrastructure/jwt"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrPersonNotFound     = errors.New("person not found")
)

type Service struct {
	JwtService *appjwt.JwtService
}

func NewAuthService(jwtService *appjwt.JwtService) *Service {
	return &Service{
		JwtService: jwtService,
	}
}
