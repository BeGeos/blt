package jwt

import (
	"go-hex/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct{}

func (s *JwtService) NewAccessToken(c AccessTokenClaims) (string, error) {
	cfg, _ := config.Load()
	claims := &JwtClaims{
		AccessTokenClaims: AccessTokenClaims{
			UserID: c.UserID,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.Authentication.SigningKey))
}

func (s *JwtService) NewRefreshToken(c RefreshTokenClaims) (string, error) {
	cfg, _ := config.Load()
	claims := &RefreshJwtClaims{
		RefreshTokenClaims: RefreshTokenClaims{
			UserID:  c.UserID,
			Version: c.Version,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 30 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.Authentication.SigningKey))
}

func NewJwtService() *JwtService {
	return &JwtService{}
}
