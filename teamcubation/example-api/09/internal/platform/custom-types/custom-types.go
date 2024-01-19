package ctypes

// Error messages constants.
const (
	ErrItemNotFound     = "Item not found"    // Message for item not found error
	ErrInvalidParameter = "Invalid parameter" // Message for invalid parameter error
)

// CustomError is a custom error type that implements the error interface.
type CustomError struct {
	Code    int    // Error code
	Message string // Error message
}

// Error implements the error interface, returning the error message.
func (e *CustomError) Error() string {
	return e.Message
}

// NewCustomError creates a new instance of CustomError with the provided code and message.
func NewCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}
