package cmsapi

import (
	"fmt"
	"net/http"
)

type APIError struct {
	StatusCode int    `json:"code"`
	Type       string `json:"type,omitempty"`
	Method     string `json:"method,omitempty"`
	MetPak     string `json:"package,omitempty"`
	Message    string `json:"message,omitempty"`
}

func NewAPIError(st int, tp string, mt string, pk string, mg error) *APIError {
	return &APIError{
		StatusCode: st,
		Type:       tp,
		Method:     mt,
		MetPak:     pk,
		Message:    mg.Error(),
	}
}

func (a *APIError) Error() string {
	return fmt.Sprintf("status code: %v, error type: %v - method: %v.%v - error message: %v", a.StatusCode, a.Type, a.MetPak, a.Method, a.Message)
}

func InvalidJSON(mt string, pk string, err error) *APIError {
	return &APIError{
		StatusCode: http.StatusBadRequest,
		Type:       "invalid-json",
		Method:     mt,
		MetPak:     pk,
		Message:    "Invalid or malformed JSON, " + err.Error(),
	}
}

func BadRequest(mt string, pk string, err error) *APIError {
	return &APIError{
		StatusCode: http.StatusBadRequest,
		Type:       "client-error",
		Method:     mt,
		MetPak:     pk,
		Message:    "Cannot process current request, " + err.Error(),
	}
}

func InternalServerError(mt string, pk string, err error) *APIError {
	return &APIError{
		StatusCode: http.StatusInternalServerError,
		Type:       "server-error",
		Method:     mt,
		MetPak:     pk,
		Message:    "Internal server error, " + err.Error(),
	}
}
