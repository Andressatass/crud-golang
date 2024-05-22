package handler

import (
	"crud-golang/server/viewmodel"
	"crud-golang/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HandleNewUserAddres(
	w http.ResponseWriter,
	r *http.Request,
	userAddresService service.UserAddresService) {
	var userAddresModel viewmodel.UserAddresModel
	var errors []error

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		WriteErrorResponse(w, err, http.StatusOK)

		return
	}

	err = json.Unmarshal([]byte(bytes), &userAddresModel)
	if err != nil {
		WriteErrorResponse(w, err, http.StatusOK)

		return
	}

	errors = userAddresModel.Validate()
	if len(errors) != 0 {
		WriteErrorResponse(w, errors[0], http.StatusBadRequest)

		return
	}

	domainUserAddres := userAddresModel.Parse()

	id, err := userAddresService.NewUserAddres(domainUserAddres)
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

func HandleFindUserAddres(
	w http.ResponseWriter,
	r *http.Request,
	userAddresService service.UserAddresService) {
	usersAddres, err := userAddresService.FindUsersAddres()
	if err != nil {
		WriteErrorResponse(w, err, http.StatusOK)

		return
	}

	byte, err := json.Marshal(usersAddres)
	if err != nil {
		fmt.Fprintf(w, "Erro ao codificar o erro: %v", err)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byte)
}

func HandleUpdateUserAddres(
	w http.ResponseWriter,
	r *http.Request,
	userAddresService service.UserAddresService) {
	var userAddresModel viewmodel.UserAddresModel
	var errors []error

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		WriteErrorResponse(w, err, http.StatusOK)

		return
	}

	err = json.Unmarshal([]byte(bytes), &userAddresModel)
	if err != nil {
		WriteErrorResponse(w, err, http.StatusOK)

		return
	}

	errors = userAddresModel.Validate()
	if len(errors) != 0 {
		WriteErrorResponse(w, errors[0], http.StatusBadRequest)

		return
	}

	domainUserAddres := userAddresModel.Parse()

	err = userAddresService.UpdateUserAddres(domainUserAddres)
	if err != nil {
		WriteErrorResponse(w, err, http.StatusOK)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func HandleDeleteUserAddres(
	w http.ResponseWriter,
	r *http.Request,
	userAddresService service.UserAddresService) {
	var userAddresModel viewmodel.UserAddresModel

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		WriteErrorResponse(w, err, http.StatusOK)

		return
	}

	err = json.Unmarshal([]byte(bytes), &userAddresModel)
	if err != nil {
		WriteErrorResponse(w, err, http.StatusOK)

		return
	}

	domainUserAddres := userAddresModel.Parse()

	err = userAddresService.DeleteUserAddres(domainUserAddres)
	if err != nil {
		WriteErrorResponse(w, err, http.StatusOK)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
