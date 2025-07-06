package auth

type RefreshTokenRequestDto struct {
	Token string `json:"token" validate:"required"`
}

type RegisterRequestDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type ResetPasswordRequestDto struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordDto struct {
	Password string `json:"password" validate:"required,min=6"`
	Token    string `param:"token"`
}
