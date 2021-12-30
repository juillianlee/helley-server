package contract

type (
	LoginRequest struct {
		Username string
		Password string
	}

	RefreshTokenRequest struct {
		RefreshToken string
	}
)
