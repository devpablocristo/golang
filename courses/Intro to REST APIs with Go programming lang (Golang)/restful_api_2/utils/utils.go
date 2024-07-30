package utils

import (
	"encoding/json"
	"net/http"
	"restful_api_2/models"
)

func EnviarError(w http.ResponseWriter, status int, err models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

func EnviarExito(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}
