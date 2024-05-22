package entities

import (
	"gorm.io/gorm"
)

// User holds user information from database.
type User struct {
	gorm.Model
	Id       string `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Age      int    `gorm:"not null"`
	AddresID uint   `gorm:"foreignKey"`
	//Addres   UserAddres `gorm:"foreignKey:AddressID"`
}
