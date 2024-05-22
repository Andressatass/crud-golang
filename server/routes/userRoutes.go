package routes

import (
	httpHandler "crud-golang/server/handler"
	"crud-golang/service"
	"net/http"
)

func Register(userService service.UserService) {

	http.HandleFunc("/NewUser", func(w http.ResponseWriter, r *http.Request) {
		httpHandler.HandleNewUser(w, r, userService)
	})

}
