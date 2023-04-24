package presenter

type ItemResponse struct {
	Error bool     `json:"error"`
	Data  jsonItem `json:"data"`
}

type ItemsResponse struct {
	Error bool       `json:"error"`
	Data  []jsonItem `json:"data"`
}
