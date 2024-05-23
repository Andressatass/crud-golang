package entities

import (
	"gorm.io/gorm"
)

// User holds user information from database.
type User struct {
	gorm.Model
	Name string `gorm:"not null"`
	Age  int    `gorm:"not null"`
}
