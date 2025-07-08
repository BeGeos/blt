package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type GetConfigArgs struct {
	authenticationRequired bool
}

type AccessTokenClaims struct {
	UserID uint `json:"user_id"`
}

type RefreshTokenClaims struct {
	UserID  uint `json:"user_id"`
	Version uint `json:"version"`
}

type JwtClaims struct {
	AccessTokenClaims
	jwt.RegisteredClaims
}

type RefreshJwtClaims struct {
	RefreshTokenClaims
	jwt.RegisteredClaims
}
