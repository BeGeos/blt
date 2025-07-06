package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	appjwt "github.com/BeGeos/go-echo/cmd/apps/auth/jwt"
)

type Services struct {
	// Injected services
	Token    Token
	Password Password
}

func NewServices() *Services {
	return &Services{
		Token:    newAuthTokenService(),
		Password: newPasswordService(),
	}
}

type Token interface {
	GetValidTokens(userID, version uint) (Tokens, error)
	GetNewAccessToken(userID uint) (string, error)
}

// Subservices - this are private
type _AuthTokenService struct {
	Jwt appjwt.Jwt
}

func newAuthTokenService() Token {
	return &_AuthTokenService{
		Jwt: appjwt.NewJwt(),
	}
}

func (s *_AuthTokenService) GetValidTokens(userID, version uint) (Tokens, error) {
	type ResultCh struct {
		token string
		err   error
	}

	accessCh := make(chan ResultCh)
	refreshCh := make(chan ResultCh)

	go func(ch chan ResultCh) {
		claims := appjwt.AccessTokenClaims{
			UserID: userID,
		}
		token, err := s.Jwt.New(claims, appjwt.NewJwtArgs{Kind: "access"})
		ch <- ResultCh{token: token, err: err}
	}(accessCh)

	go func(ch chan ResultCh) {
		claims := appjwt.RefreshTokenClaims{
			UserID:  userID,
			Version: version, // replace with actual version check logic
		}
		token, err := s.Jwt.New(claims, appjwt.NewJwtArgs{Kind: "refresh"})
		ch <- ResultCh{token: token, err: err}
	}(refreshCh)

	accessResult := <-accessCh
	refreshResult := <-refreshCh

	if accessResult.err != nil || refreshResult.err != nil {
		return Tokens{}, errors.New("failed to create token")
	}

	return Tokens{
		AccessToken:  accessResult.token,
		RefreshToken: refreshResult.token,
	}, nil
}

func (s *_AuthTokenService) GetNewAccessToken(userID uint) (string, error) {
	claims := appjwt.AccessTokenClaims{
		UserID: userID,
	}

	token, err := s.Jwt.New(claims, appjwt.NewJwtArgs{Kind: "access"})
	if err != nil {
		return "", errors.New("failed to create token")
	}

	return token, nil
}

type Password interface {
	Hash(password string) ([]byte, error)
	Compare(hashedPassword, password string) error
}
type _PasswordService struct{}

func newPasswordService() Password {
	return &_PasswordService{}
}

type HashPasswordArgs struct {
	Cost int
}

func (s *_PasswordService) Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (s *_PasswordService) Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
