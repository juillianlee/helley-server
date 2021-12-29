package login

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type LoginUseCase interface {
	Handle(username string, password string) (map[string]string, error)
}

type loginUseCase struct {
}

func NewLoginUseCase() LoginUseCase {
	return &loginUseCase{}
}

func (h *loginUseCase) Handle(username string, password string) (map[string]string, error) {

	if username == "juillian" && password == "abc123" {
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Juillian Lee"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return map[string]string{}, err
		}

		return map[string]string{
			"token": t,
		}, nil

	}

	return nil, errors.New("ErrUnauthorized")
}
