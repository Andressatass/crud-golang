package routes

import (
	"crud-golang/server/handler"
	"crud-golang/service"
	"net/http"
)

func RegisterUserRoutes(userService service.UserService) {

	http.HandleFunc("/NewUser", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleNewUser(w, r, userService)
	})
	http.HandleFunc("/FindAllUsers", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleFindUsers(w, r, userService)
	})
	http.HandleFunc("/UpdateUser", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleUpdateUser(w, r, userService)
	})
	http.HandleFunc("/DeleteUser", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleDeleteUser(w, r, userService)
	})

}
