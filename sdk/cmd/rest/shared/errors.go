package event

import (
	"fmt"
	"log"
)

type CustomError struct {
	Message string
	Err     error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func New(msg string, err error) error {
	return &CustomError{
		Message: msg,
		Err:     err,
	}
}

func HandleFatalError(msg string, err error) {
	log.Fatalf("%s: %v", msg, err)
}

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewAPIError(code int, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
}

type APIMessage struct {
	Message string      `json:"message,omitempty"`
	Item    interface{} `json:"item,omitempty"`
}

func NewAPIMessage(message string, items ...interface{}) *APIMessage {
	apiMessage := &APIMessage{
		Message: message,
	}
	if len(items) > 0 {
		apiMessage.Item = items[0]
	}
	return apiMessage
}

// type APISliceMessage struct {
// 	Message string      `json:"message,omitempty"`
// 	Items   interface{} `json:"items"`
// }

// func NewAPISliceMessage(message string, items ...[]interface{}) *APISliceMessage {
// 	return &APISliceMessage{
// 		Message: message,
// 		Items:   items,
// 	}
// }
