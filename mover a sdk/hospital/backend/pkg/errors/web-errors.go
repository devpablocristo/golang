package customErr

import (
	"errors"
	"fmt"
)

var mapErrors = map[int]string{400: "param name cannot be empty"}

type RequestError struct {
	StatusCode int

	Err error
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

func DoRequest(code int) error {

	return &RequestError{
		StatusCode: code,
		Err:        errors.New(mapErrors[code]),
	}
}
