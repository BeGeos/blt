package auth_services

import (
	"time"

	"github.com/BeGeos/go-echo/internal/settings"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(settings.JwtSecretKey)

type PersonClaims struct {
	UserID  uint `json:"user_id"`
	Version uint `json:"version"`
}
type JwtClaims struct {
	PersonClaims
	jwt.RegisteredClaims
}

type JwtService struct{}

func (s *JwtService) NewToken(c PersonClaims) (string, error) {
	claims := &JwtClaims{
		PersonClaims: PersonClaims{
			UserID:  c.UserID,
			Version: c.Version,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
