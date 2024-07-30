package domain

import (
	"time"

	pdomain "github.com/devpablocristo/golang/06-apps/qh/person/domain"
)

type Book struct {
	Author pdomain.Person `json:"author"`
	Title  string         `json:"title"`
	Price  float64        `json:"price"`
	ISBN   string         `json:"isbn"`
}

type BookStock struct {
	Book      *Book     `json:"book"`
	Stock     int64     `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
}
