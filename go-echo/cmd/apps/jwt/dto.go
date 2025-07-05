package jwt

type RefreshTokenRequestDto struct {
	Token string `json:"token" validate:"required"`
}
