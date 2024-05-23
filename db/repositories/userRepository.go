package repositories

import (
	"crud-golang/db/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r UserRepository) Create(user entities.User) (uint, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

func (r UserRepository) FindAllUsers() ([]entities.User, error) {
	var users []entities.User

	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (r UserRepository) Update(user entities.User) error {
	result := r.db.Model(&user).Where("uuid = ?", user.UUID).Updates(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r UserRepository) Delete(user entities.User) error {
	result := r.db.Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
