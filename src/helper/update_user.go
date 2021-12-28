package helper

type UpdateUserRequest struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}
