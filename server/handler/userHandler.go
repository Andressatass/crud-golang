package handler

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
		WriteErrorResponse(w, err, http.StatusOK)

		return
	}

	err = json.Unmarshal([]byte(bytes), &userModel)
	if err != nil {
		WriteErrorResponse(w, err, http.StatusOK)

		return
	}

	errors = userModel.Validate()
	if len(errors) != 0 {
		WriteErrorResponse(w, errors[0], http.StatusBadRequest)

		return
	}

	domainUser := userModel.Parse()

	id, err := userService.NewUser(domainUser)
	if err != nil {
		WriteErrorResponse(w, err, http.StatusOK)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	byte, err := json.Marshal(id)
	if err != nil {
		fmt.Fprintf(w, "Erro ao codificar o erro: %v", err)

		return
	}
	w.Write(byte)
}

// talvez editar o modelo que é devolvido, para n ir com data de criação etc.
func HandleFindUsers(w http.ResponseWriter, r *http.Request, userService service.UserService) {
	users, err := userService.FindUsers()
	if err != nil {
		WriteErrorResponse(w, err, http.StatusOK)

		return
	}

	byte, err := json.Marshal(users)
	if err != nil {
		fmt.Fprintf(w, "Erro ao codificar o erro: %v", err)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byte)
}
