package helper

type StoreUserRequest struct {
	Name     string
	Email    string
	Password string
}

type StoreUserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
