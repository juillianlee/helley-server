package app_repository

import (
	"app-helley/src/domain"
)

type UserRepository interface {
	Store(user domain.User) (domain.User, error)
	DeleteById(id string) error
	FindById(id string) (domain.User, error)
	Update(user domain.User) error
	Find() ([]domain.User, error)
}
