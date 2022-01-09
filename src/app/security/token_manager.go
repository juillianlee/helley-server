package security

import (
	"errors"
)

var ErrUnauthorized = errors.New("unauthorized")
var ErrForbidenAcess = errors.New("forbiden")
var ErrUnexpectedSignin = errors.New("unexpected signing")

type TokenManager interface {
	GenerateTokenPair() (TokenPayload, error)
	RefreshToken(refreshToken string) (TokenPayload, error)
}
type TokenPayload struct {
	AccessToken  string
	RefreshToken string
}
