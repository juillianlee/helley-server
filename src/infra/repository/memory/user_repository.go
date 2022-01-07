package repository_memory

import (
	app_repository "app-helley/src/app/repository"
	domain_user "app-helley/src/domain/user"
)

type UserRepository struct {
	users []domain_user.User
}

func (rep *UserRepository) Store(user domain_user.User) (domain_user.User, error) {
	rep.users = append(rep.users, user)
	return user, nil
}

func (rep *UserRepository) DeleteById(id string) error {
	for index, user := range rep.users {
		if user.ID == id {
			rep.users = append(rep.users[:index], rep.users[index+1:]...)
			return nil
		}
	}
	return app_repository.ErrNotFoundRegister
}

func (rep *UserRepository) FindById(id string) (domain_user.User, error) {
	for _, user := range rep.users {
		if user.ID == id {
			return user, nil
		}
	}
	return domain_user.User{}, app_repository.ErrNotFoundRegister
}

func (rep *UserRepository) Update(userUpdate domain_user.User) error {
	for _, user := range rep.users {
		if user.ID == userUpdate.ID {
			user.Name = userUpdate.Name
			user.Email = userUpdate.Email
			user.CreatedAt = userUpdate.CreatedAt
			user.UpdateAt = userUpdate.UpdateAt
			return nil
		}
	}
	return app_repository.ErrNotFoundRegister
}

func (rep *UserRepository) Find() ([]domain_user.User, error) {
	return rep.users, nil
}

func (rep *UserRepository) FindByEmail(email string) (domain_user.User, error) {
	for _, user := range rep.users {
		if user.Email == email {
			return user, nil
		}
	}
	return domain_user.User{}, app_repository.ErrNotFoundRegister
}
