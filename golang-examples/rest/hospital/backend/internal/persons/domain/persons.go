package domain

type Person struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	DNI      int64  `json:"dni"`
	Gender   string `json:"gender"`
}
