package viewmodel

import (
	"crud-golang/domain"
	"fmt"
)

type UserModel struct {
	Id   string `json:"user_uuid"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u UserModel) Validate() []error {
	genericErrorFormat := "invalid-parameter-%s"
	var Errors []error

	if u.Age <= 0 {
		Errors = append(Errors, fmt.Errorf(genericErrorFormat, "age"))
	}

	if u.Name == "" {
		Errors = append(Errors, fmt.Errorf(genericErrorFormat, "name"))
	}

	if len(Errors) > 0 {
		return Errors
	}

	return nil
}

func (u UserModel) Parse() domain.User {
	return domain.User{
		Id:   u.Id,
		Name: u.Name,
		Age:  u.Age,
	}
}
