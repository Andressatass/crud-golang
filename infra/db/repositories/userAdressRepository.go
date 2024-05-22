package repositories

import (
	"crud-golang/infra/db/entities"

	"gorm.io/gorm"
)

type userAdressRepository struct {
	db *gorm.DB
}

func NewUserAdressRepository(db *gorm.DB) userAdressRepository {
	return userAdressRepository{
		db: db,
	}
}

func (r userAdressRepository) Create(userAdress entities.UserAddres) (uint, error) {
	result := r.db.Create(&userAdress)
	if result.Error != nil {
		return 0, result.Error
	}

	return userAdress.ID, nil
}

func (r userAdressRepository) FindAllUsers() ([]entities.UserAddres, error) {
	var usersAdress []entities.UserAddres

	result := r.db.Find(&usersAdress)
	if result.Error != nil {
		return nil, result.Error
	}

	return usersAdress, nil
}

func (r userAdressRepository) Update(userAddres entities.UserAddres) error {
	result := r.db.Save(userAddres)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r userAdressRepository) Delete(userAddres entities.UserAddres) error {
	result := r.db.Delete(&userAddres)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
