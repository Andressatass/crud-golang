package domain

import "crud-golang/infra/db/entities"

//User holds the user information
type User struct {
	Id   string
	Name string
	Age  int
}

func (u User) ParseToEntity() entities.User {
	return entities.User{
		Id:   u.Id,
		Name: u.Name,
		Age:  u.Age,
	}
}
