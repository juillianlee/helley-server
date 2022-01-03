package presentation

type (
	UserResponse struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	StoreUserRequest struct {
		Name     string `validate:"required"`
		Email    string `validate:"required,email"`
		Password string `validate:"required"`
	}

	StoreUserResponse struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	UpdateUserRequest struct {
		Name  string `validate:"required"`
		Email string `validate:"required,email"`
	}
)
