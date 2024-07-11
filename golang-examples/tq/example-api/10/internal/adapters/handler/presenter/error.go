package presenter

// ApiError defines the structure for API error responses.
type ApiError struct {
	StatusCode int    `json:"status_code"` // HTTP status code
	Message    string `json:"message"`     // Error message
}

// NewApiError creates a new ApiError with the given status code and message.
func NewApiError(statusCode int, message string) *ApiError {
	return &ApiError{
		StatusCode: statusCode,
		Message:    message,
	}
}
