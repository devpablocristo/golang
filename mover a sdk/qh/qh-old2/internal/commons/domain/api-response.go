package cmsdomain

import (
	"encoding/json"
	"net/http"
)

// eliminar
type ResponseInfo struct {
	Status int `json:"status"`
	Data   any `json:"data"`
}

type ResponseAPI struct {
	Success bool `json:"success"`
	Status  int  `json:"status,omitempty"`
	Result  any  `json:"result,omitempty"`
}

func Success(result any, status int) *ResponseAPI {
	return &ResponseAPI{
		Success: true,
		Status:  status,
		Result:  result,
	}
}

func (r *ResponseAPI) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	return json.NewEncoder(w).Encode(r)
}
