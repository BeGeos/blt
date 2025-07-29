package jwt

import (
	"go-hex/internal/apps/apperrors"
	"go-hex/internal/config"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct{}

var (
	ErrInvalidTokenVersion = apperrors.New("invalid_token_version", "token is invalid or expired", apperrors.WithHttpCode(http.StatusUnauthorized))
	ErrFailedToken         = apperrors.New("generic_error", "failed to create token", apperrors.WithHttpCode(http.StatusInternalServerError))
)

func (s *JwtService) NewAccessToken(userID int) (string, error) {
	cfg, _ := config.Load()
	claims := &JwtClaims{
		Typ: "access",
		Sub: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.Authentication.SigningKey))
}

func (s *JwtService) NewRefreshToken(userID, version int) (string, error) {
	cfg, _ := config.Load()
	claims := &RefreshJwtClaims{
		Version: version,
		Typ:     "refresh",
		Sub:     userID,
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
