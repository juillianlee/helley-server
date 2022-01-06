package domain_user

import (
	"app-helley/src/app/validator"
	"time"
)

type User struct {
	ID        string
	Name      string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string
	CreatedAt time.Time
	UpdateAt  time.Time
}

func (u *User) Validate() error {
	return validator.Validate(u)
}
