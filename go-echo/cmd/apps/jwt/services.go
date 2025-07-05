package jwt

import (
	"time"

	"github.com/BeGeos/go-echo/internal/settings"
	"github.com/golang-jwt/jwt/v5"
)

type Token interface {
	New(claims interface{}, args NewArgs) (string, error)
}
type (
	token    struct{}
	Services struct {
		Token Token
	}
)

var jwtSecret = []byte(settings.JwtSecretKey)

type NewArgs struct {
	Kind string
}

func (s *token) New(claims interface{}, args NewArgs) (string, error) {
	switch args.Kind {
	case "access":
		claims, _ := claims.(AccessTokenClaims)
		return s.newAccessToken(claims)
	case "refresh":
		claims, _ := claims.(RefreshTokenClaims)
		return s.newRefreshToken(claims)
	}

	return "", jwt.ErrInvalidKeyType
}

func (s *token) newAccessToken(c AccessTokenClaims) (string, error) {
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

func (s *token) newRefreshToken(c RefreshTokenClaims) (string, error) {
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

func newTokenService() Token {
	return &token{}
}

func NewServices() *Services {
	return &Services{
		Token: newTokenService(),
	}
}
