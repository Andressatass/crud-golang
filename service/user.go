package service

import (
	"crud-golang/db/entities"
	"crud-golang/db/repositories"
	"crud-golang/domain"
	"fmt"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func (us UserService) NewUser(user domain.User) (uint, error) {
	userEntity := user.ParseToEntity()

	uui, err := us.UserRepository.Create(userEntity)
	if err != nil {
		return 0, fmt.Errorf("could not create user:%w", err)
	}

	return uui, nil
}

func (us UserService) FindUsers() ([]entities.User, error) {
	var users []entities.User

	users, err := us.UserRepository.FindAllUsers()
	if err != nil {
		return users, fmt.Errorf("could not find users: %w", err)
	}

	return users, nil
}

func (us UserService) UpdateUser(user domain.User) error {
	userEntity := user.ParseToEntity()

	err := us.UserRepository.Update(userEntity)
	if err != nil {
		return fmt.Errorf("could not update user: %w", err)
	}

	return nil
}

func (us UserService) DeleteUser(user domain.User) error {
	userEntity := user.ParseToEntity()

	err := us.UserRepository.Delete(userEntity)
	if err != nil {
		return fmt.Errorf("could not delete user:%w", err)
	}

	return nil
}
