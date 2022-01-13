package dto

type LoginRequest struct {
	Username string `validate:"required,email"`
	Password string `validate:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `validate:"required"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type CreateAccountRequest struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}
