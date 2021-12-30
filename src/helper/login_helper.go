package helper

type (
	LoginRequest struct {
		Username string
		Password string
	}

	RefreshTokenRequest struct {
		RefreshToken string
	}
)
