package entities

import "gorm.io/gorm"

// UserAddres holds the user addres information from database.
type UserAddres struct {
	gorm.Model
	Street string `gorm:"not null"`
	City   string `gorm:"not null"`
	State  string `gorm:"not null"`
}
