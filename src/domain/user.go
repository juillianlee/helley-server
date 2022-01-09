package domain

import (
	"app-helley/src/app/validator"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string
	Name      string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string
	CreatedAt time.Time
	UpdateAt  time.Time
}

// Compra a hash da senha com senha do usuario
func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// Valida a estrutura do usuario conforme definido no dominio
func (u *User) Validate() error {
	return validator.Validate(u)
}

// Gera uma hash da senha utilizando o bcrypt
func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
