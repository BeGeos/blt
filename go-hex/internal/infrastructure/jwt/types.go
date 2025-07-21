package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type GetConfigArgs struct {
	AuthenticationRequired bool
}

type JwtClaims struct {
	Typ string `json:"typ,omitempty"` // Type of the token, e.g., "access"
	Sub int    `json:"sub,omitempty"` // Subject, typically the user ID
	jwt.RegisteredClaims
}

type RefreshJwtClaims struct {
	Version int    `json:"version"`
	Typ     string `json:"typ,omitempty"`
	Sub     int    `json:"sub,omitempty"`
	jwt.RegisteredClaims
}
