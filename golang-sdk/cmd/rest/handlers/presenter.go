package handler

type CreateEventResponse struct {
	Message string `json:"message"`
	Err     string `json:"err,omitempty"`
}
