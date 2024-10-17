package shared

import (
	"fmt"
)

// WrappedError es un tipo de error personalizado que encapsula un mensaje y un error original.
type WrappedError struct {
	Message string
	Err     error
}

// Error implementa la interfaz de error para WrappedError.
func (e *WrappedError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// NewCustomError crea un nuevo WrappedError.
func NewCustomError(msg string, err error) *WrappedError {
	return &WrappedError{
		Message: msg,
		Err:     err,
	}
}
