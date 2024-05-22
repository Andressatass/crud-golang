package repositories

import (
	"crud-golang/infra/db/entities"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) userRepository {
	return userRepository{
		db: db,
	}
}

func (r userRepository) Create(user entities.User) (uint, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

func (r userRepository) FindAllUsers() ([]entities.User, error) {
	var users []entities.User

	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (r userRepository) Update(user entities.User) error {
	result := r.db.Save(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r userRepository) Delete(user entities.User) error {
	result := r.db.Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
