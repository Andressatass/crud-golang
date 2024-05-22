package service

import (
	"crud-golang/domain"
	"crud-golang/infra/db/repositories"
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

func (us UserService) DeleteUser() {

}
