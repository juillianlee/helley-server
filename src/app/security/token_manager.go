package security

import (
	"app-helley/src/domain"
	"errors"
)

var ErrUnauthorized = errors.New("unauthorized")
var ErrForbidenAcess = errors.New("forbiden")
var ErrUnexpectedSignin = errors.New("unexpected signing")
var ErrUserNameOrPasswordInvalid = errors.New("username or password invalid")

type TokenManager interface {
	GenerateTokenPair(u domain.User) (TokenPayload, error)
	ValidateRefreshToken(refreshToken string) (map[string]interface{}, error)
	ValidateAccessToken(accessToken string) (map[string]interface{}, error)
}
type TokenPayload struct {
	AccessToken  string
	RefreshToken string
}
