package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type GetConfigArgs struct {
	AuthenticationRequired bool
}

type AccessTokenClaims struct {
	UserID int `json:"user_id"`
}

type RefreshTokenClaims struct {
	UserID  int `json:"user_id"`
	Version int `json:"version"`
}

type JwtClaims struct {
	AccessTokenClaims
	jwt.RegisteredClaims
}

type RefreshJwtClaims struct {
	RefreshTokenClaims
	jwt.RegisteredClaims
}
