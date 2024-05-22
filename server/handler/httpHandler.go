package httpHandler

import (
	"crud-golang/server/viewmodel"
	"crud-golang/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HandleNewUser(w http.ResponseWriter, r *http.Request, userService service.UserService) {
	var userModel viewmodel.UserModel
	var errors []error

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(bytes), &userModel)
	if err != nil {
		return
	}

	errors = userModel.Validate()
	if len(errors) != 0 {
		return
	}

	domainUser := userModel.Parse()

	id, err := userService.NewUser(domainUser)
	if err != nil {
		return
	}

	fmt.Print(id)

}
