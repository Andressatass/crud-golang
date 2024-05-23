package domain

import "crud-golang/db/entities"

//User holds the user information
type User struct {
	UUID string
	Name string
	Age  int
}

func (u User) ParseToEntity() entities.User {
	return entities.User{
		UUID: u.UUID,
		Name: u.Name,
		Age:  u.Age,
	}
}
