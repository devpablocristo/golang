package cmsapi

type APIResponse struct {
	Success bool `json:"success"`
	Status  int  `json:"status,omitempty"`
	Result  any  `json:"result,omitempty"`
}

func NewAPIRespone(sc bool, st int, rt any) *APIResponse {
	return &APIResponse{
		Success: sc,
		Status:  st,
		Result:  rt,
	}
}

func SuccessResponse(status int, result any) *APIResponse {
	return &APIResponse{
		Success: true,
		Status:  status,
		Result:  result,
	}
}

func FailResponse(err *APIError) *APIResponse {
	return &APIResponse{
		Success: false,
		Status:  err.StatusCode,
		Result:  err.Message,
	}
}
