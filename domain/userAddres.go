package domain

import "crud-golang/db/entities"

//UserAddres holds the user addres information
type UserAddres struct {
	Street string
	City   string
	State  string
}

func (ua UserAddres) ParseToEntity() entities.UserAddres {
	return entities.UserAddres{
		Street: ua.Street,
		City:   ua.City,
		State:  ua.State,
	}
}
