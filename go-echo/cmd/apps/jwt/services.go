package jwt

import (
	"time"

	"github.com/BeGeos/go-echo/internal/settings"
	"github.com/golang-jwt/jwt/v5"
)

type Jwt interface {
	New(claims interface{}, args NewArgs) (string, error)
}
type (
	Services struct {
		// Injected
		Jwt Jwt
	}
	token struct{}
)

var jwtSecret = []byte(settings.JwtSecretKey)

type NewArgs struct {
	Kind string
}

func (s *token) New(claims interface{}, args NewArgs) (string, error) {
	switch args.Kind {
	case "access":
		claims, _ := claims.(AccessTokenClaims)
		return newAccessToken(claims)
	case "refresh":
		claims, _ := claims.(RefreshTokenClaims)
		return newRefreshToken(claims)
	}

	return "", jwt.ErrInvalidKeyType
}

func newAccessToken(c AccessTokenClaims) (string, error) {
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

func newRefreshToken(c RefreshTokenClaims) (string, error) {
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

func newJwtService() Jwt {
	return &token{}
}

func NewServices() *Services {
	return &Services{
		Jwt: newJwtService(),
	}
}
