package dto

type (
	LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required"`
	}

	RefreshTokenRequest struct {
		RefreshToken string `validate:"required"`
	}

	TokenResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)
