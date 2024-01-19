package ctypes

const (
	ErrItemNotFound     = "Item not found"
	ErrInvalidParameter = "Invalid parameter"
)

// CustomError es un tipo de error personalizado que implementa la interfaz error.
type CustomError struct {
	Code    int
	Message string
}

// Error implementa la interfaz error, devolviendo el mensaje del error.
func (e *CustomError) Error() string {
	return e.Message
}

// NewCustomError crea una nueva instancia de CustomError con el código y mensaje proporcionados.
func NewCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}
