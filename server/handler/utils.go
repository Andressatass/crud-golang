package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func WriteErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	errorResponse := ErrorResponse{
		Message: err.Error(),
		Code:    statusCode,
	}

	jsonData, err := json.Marshal(errorResponse)
	if err != nil {
		fmt.Fprintf(w, "Erro ao codificar o erro: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
