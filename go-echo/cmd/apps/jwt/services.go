package jwt

import (
	"time"

	"github.com/BeGeos/go-echo/internal/settings"
	"github.com/golang-jwt/jwt/v5"
)

type _Services struct{}

var jwtSecret = []byte(settings.JwtSecretKey)

func (s *_Services) NewAccessToken(c AccessTokenClaims) (string, error) {
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
	return token.SignedString(jwtSecret)
}

func (s *_Services) NewRefreshToken(c RefreshTokenClaims) (string, error) {
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
	return token.SignedString(jwtSecret)
}

var Services = &_Services{}
