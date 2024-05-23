package repositories

import (
	"crud-golang/db/entities"

	"gorm.io/gorm"
)

type UserAdressRepository struct {
	db *gorm.DB
}

func NewUserAdressRepository(db *gorm.DB) UserAdressRepository {
	return UserAdressRepository{
		db: db,
	}
}

func (r UserAdressRepository) Create(userAdress entities.UserAddres) (uint, error) {
	result := r.db.Create(&userAdress)
	if result.Error != nil {
		return 0, result.Error
	}

	return userAdress.ID, nil
}

func (r UserAdressRepository) FindAllUsers() ([]entities.UserAddres, error) {
	var usersAdress []entities.UserAddres

	result := r.db.Find(&usersAdress)
	if result.Error != nil {
		return nil, result.Error
	}

	return usersAdress, nil
}

func (r UserAdressRepository) Update(userAddres entities.UserAddres) error {
	result := r.db.Model(&userAddres).Where("uuid = ?", userAddres.UUID).Updates(userAddres)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r UserAdressRepository) Delete(userAddres entities.UserAddres) error {
	result := r.db.Delete(&userAddres)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
