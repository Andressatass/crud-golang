package main

import (
	"crud-golang/configuration"
	"crud-golang/infra/db/repositories"
	"crud-golang/infra/db/settings"
	"crud-golang/server"
	"crud-golang/server/routes"
	"crud-golang/service"
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

	userService := service.UserService{
		UserRepository: userRepository,
	}

	routes.Register(userService)

	err = server.StartServer(":8080")
	if err != nil {
		return
	}

}
