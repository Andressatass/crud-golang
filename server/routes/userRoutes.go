package routes

import (
	"crud-golang/server/handler"
	"crud-golang/service"
	"net/http"
)

func Register(userService service.UserService) {

	http.HandleFunc("/NewUser", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleNewUser(w, r, userService)
	})
	http.HandleFunc("/FindAllUsers", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleFindUsers(w, r, userService)
	})
}
