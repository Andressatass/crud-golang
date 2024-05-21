package entities

import "gorm.io/gorm"

// User holds user information from database.
type User struct {
	gorm.Model
	Id   string
	Name string
	Age  int
}
