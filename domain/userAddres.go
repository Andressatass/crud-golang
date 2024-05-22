package domain

import "crud-golang/db/entities"

//UserAddres holds the user addres information
type UserAddres struct {
	Id     string
	Street string
	City   string
	State  string
	//User  User
}

func (ua UserAddres) ParseToEntity() entities.UserAddres {
	return entities.UserAddres{
		Id:     ua.Id,
		Street: ua.Street,
		City:   ua.City,
		State:  ua.State,
	}
}
