package entities

import "gorm.io/gorm"

// UserAddres holds the user addres information from database.
type UserAddres struct {
	gorm.Model
	Id     string
	Street string
	City   string
	State  string
}
