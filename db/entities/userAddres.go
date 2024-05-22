package entities

import "gorm.io/gorm"

// UserAddres holds the user addres information from database.
type UserAddres struct {
	gorm.Model
	Id     string `gorm:"primaryKey"`
	Street string `gorm:"not null"`
	City   string `gorm:"not null"`
	State  string `gorm:"not null"`
	//User   User   `gorm:"foreignKey:UserID"`
}
