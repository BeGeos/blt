package auth

import (
	"go-hex/internal/apps/auth"
	"go-hex/internal/apps/core"
	appjwt "go-hex/internal/infrastructure/jwt"
)

type Service struct {
	JwtService *appjwt.JwtService
}

func NewAuthService(jwtService *appjwt.JwtService) *Service {
	return &Service{
		JwtService: jwtService,
	}
}

func (s *Service) Login(email, password string) (auth.Tokens, *core.Error) {
	// TODO: set proper user validation logic
	const op = "auth.Login"

	if email == "admin@mail.com" && password == "password" {
		accessToken, err := s.JwtService.NewAccessToken(69)
		if err != nil {
			return auth.Tokens{}, &core.Error{
				Err: err,
				Op:  op,
			}
		}

		refreshToken, err := s.JwtService.NewRefreshToken(69, 2)
		if err != nil {
			return auth.Tokens{}, &core.Error{
				Err: err,
				Op:  op,
			}
		}

		return auth.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}, nil
	}

	return auth.Tokens{}, &core.Error{
		Code:    core.ErrInvalid,
		Message: "invalid credentials",
		Op:      op,
	}
}

func (s *Service) Refresh(userID, version int) (auth.AccessToken, *core.Error) {
	const op = "auth.Refresh"
	// check version with user version
	if version != 2 {
		return auth.AccessToken{}, &core.Error{
			Code:    core.ErrInvalid,
			Message: "invalid token version",
			Op:      op,
		}
	}

	token, err := s.JwtService.NewAccessToken(userID)
	if err != nil {
		return auth.AccessToken{}, &core.Error{
			Err: err,
			Op:  op,
		}
	}

	return auth.AccessToken{AccessToken: token}, nil
}
