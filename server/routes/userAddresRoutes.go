package routes

import (
	"crud-golang/server/handler"
	"crud-golang/service"
	"net/http"
)

func RegisterUserAddresRoutes(userAddresService service.UserAddresService) {

	http.HandleFunc("/NewUserAddres", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleNewUserAddres(w, r, userAddresService)
	})
	http.HandleFunc("/FindAllUsersAddres", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleFindUserAddres(w, r, userAddresService)
	})
	http.HandleFunc("/UpdateUserAddres", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleUpdateUserAddres(w, r, userAddresService)
	})
	http.HandleFunc("/DeleteUseraddres", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleDeleteUserAddres(w, r, userAddresService)
	})
}
