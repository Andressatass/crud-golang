package service

import (
	"crud-golang/domain"
	"crud-golang/infra/db/entities"
	"crud-golang/infra/db/repositories"
	"fmt"
)

type UserAddresService struct {
	UserAddresRepository repositories.UserAdressRepository
}

func (us UserAddresService) NewUserAddres(userAddres domain.UserAddres) (uint, error) {
	userAddresEntity := userAddres.ParseToEntity()

	uui, err := us.UserAddresRepository.Create(userAddresEntity)
	if err != nil {
		return 0, fmt.Errorf("could not create user addres:%w", err)
	}

	return uui, nil
}

func (us UserAddresService) FindUsersAddres() ([]entities.UserAddres, error) {
	var usersAddres []entities.UserAddres

	usersAddres, err := us.UserAddresRepository.FindAllUsers()
	if err != nil {
		return usersAddres, fmt.Errorf("could not find users addres: %w", err)
	}

	//add logica de transformar em json aqui

	return usersAddres, nil
}

func (us UserAddresService) UpdateUserAddres(userAddres domain.UserAddres) error {
	userAddresEntity := userAddres.ParseToEntity()

	err := us.UserAddresRepository.Update(userAddresEntity)
	if err != nil {
		return fmt.Errorf("could not update user addres: %w", err)
	}

	return nil
}

func (us UserAddresService) DeleteUserAddres(userAddres domain.UserAddres) error {
	userAddresEntity := userAddres.ParseToEntity()

	err := us.UserAddresRepository.Delete(userAddresEntity)
	if err != nil {
		return fmt.Errorf("could not delete user addres:%w", err)
	}

	return nil
}
