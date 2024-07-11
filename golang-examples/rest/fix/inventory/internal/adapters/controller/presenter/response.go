package presenter

type BookResponse struct {
	Error bool     `json:"error"`
	Data  jsonBook `json:"data"`
}

type BooksResponse struct {
	Error bool       `json:"error"`
	Data  []jsonBook `json:"data"`
}
