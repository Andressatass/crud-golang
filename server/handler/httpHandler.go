package httpHandler

import (
	"crud-golang/server/viewmodel"
	"encoding/json"
	"io"
	"net/http"
)

func HandleNewuser(w http.ResponseWriter, r *http.Request) {
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

}
