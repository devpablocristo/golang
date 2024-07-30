package domain

type Task struct {
	UUID        string `json:"UUID"`
	Type        string `json:"type"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description"`
	Category    string `json:"category,omitempty"`
}
