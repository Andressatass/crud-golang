package main

import (
	"crud-golang/configuration"
	"crud-golang/infra/db/entities"
	"crud-golang/infra/db/repositories"
	"crud-golang/infra/db/settings"
	"fmt"
)

func main() {
	config, err := configuration.GetConfig()
	if err != nil {
		return
	}

	conn, err := settings.Connect(config.Database)
	if err != nil {
		return
	}

	userRepository := repositories.NewUserRepository(conn.DB)

	id, err := userRepository.Create(
		entities.User{
			Id:   "123",
			Name: "andressa",
			Age:  25,
		},
	)
	if err != nil {
		return
	}

	fmt.Print(id)

}
