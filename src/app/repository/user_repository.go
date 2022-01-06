package app_repository

import (
	domain_user "app-helley/src/domain/user"
)

type UserRepository interface {
	Store(user domain_user.User) (domain_user.User, error)
	DeleteById(id string) error
	FindById(id string) (domain_user.User, error)
	Update(user domain_user.User) error
	Find() ([]domain_user.User, error)
}