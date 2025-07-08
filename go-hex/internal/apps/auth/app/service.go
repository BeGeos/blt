package auth

import (
	appjwt "go-hex/internal/infrastructure/jwt"
)

type AuthService struct {
	JwtService appjwt.JwtService
}
