package main

import (
	"crud-golang/configuration"
	"crud-golang/db/repositories"
	"crud-golang/db/settings"
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
	userAddresRepository := repositories.NewUserAdressRepository(conn.DB)

	userService := service.UserService{
		UserRepository: userRepository,
	}
	userAddresService := service.UserAddresService{
		UserAddresRepository: userAddresRepository,
	}

	routes.RegisterUserRoutes(userService)
	routes.RegisterUserAddresRoutes(userAddresService)

	err = server.StartServer(config.API.Port)
	if err != nil {
		return
	}
}
