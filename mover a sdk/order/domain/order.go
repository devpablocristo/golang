package domain

import (
	"time"

	bdomain "github.com/devpablocristo/golang/06-apps/bookstore/inventory/domain"
	pdomain "github.com/devpablocristo/golang/06-apps/bookstore/person/domain"
)

type Order struct {
	Customer pdomain.Person `json:"client"`
	Date     time.Time      `json:"date"`
	Details  []OrderItems   `json:"details"`
}

type OrderItems struct {
	Book     bdomain.Book `json:"books"`
	Quantity int64        `json:"quantity"`
}

var Orders []Order
