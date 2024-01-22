package entity

// LayoutsResponse from the service
type LayoutsResponse struct {
	ID       string    `json:"id"`
	Status   string    `json:"status"`
	Variants []Variant `json:"variants"`
}

// Variant nested responses
type Variant struct {
	ID      string    `json:"id"`
	Units   []string  `json:"units"`
	Support []Support `json:"support"`
	Status  string    `json:"status"`
}

type Support struct {
	Sites   []string `json:"sites"`
	Clients []Client `json:"clients"`
}

type Client struct {
	ID uint `json:"id"`
}

type ShortcutsResponse struct {
	Title        string
	Color        string
	ShortcutsIds []string
	ID           string
	SiteId       string
}
