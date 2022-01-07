package repository_memory

import (
	app_repository "app-helley/src/app/repository"
	domain_user "app-helley/src/domain/user"
)

type userRepository struct {
	users []domain_user.User
}

func NewRepositoryMemory() app_repository.UserRepository {
	return &userRepository{
		users: []domain_user.User{},
	}
}

func (rep *userRepository) Store(user domain_user.User) (domain_user.User, error) {
	rep.users = append(rep.users, user)
	return user, nil
}

func (rep *userRepository) DeleteById(id string) error {
	for index, user := range rep.users {
		if user.ID == id {
			rep.users = append(rep.users[:index], rep.users[index+1:]...)
			return nil
		}
	}
	return app_repository.ErrNotFoundRegister
}

func (rep *userRepository) FindById(id string) (domain_user.User, error) {
	for _, user := range rep.users {
		if user.ID == id {
			return user, nil
		}
	}
	return domain_user.User{}, app_repository.ErrNotFoundRegister
}

func (rep *userRepository) Update(userUpdate domain_user.User) error {
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

func (rep *userRepository) Find() ([]domain_user.User, error) {
	return rep.users, nil
}

func (rep *userRepository) FindByEmail(email string) (domain_user.User, error) {
	for _, user := range rep.users {
		if user.Email == email {
			return user, nil
		}
	}
	return domain_user.User{}, app_repository.ErrNotFoundRegister
}
