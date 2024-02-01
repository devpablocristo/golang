package domain

import (
	"time"
)

// ID defines a custom type for item identifiers.
type ID uint

// MapRepo is a map-based repository for storing items.
type MapRepo map[ID]*Item

// Item represents an item entity with various attributes.
type Item struct {
	Code        string    // Unique code for the item
	Title       string    // Title of the item
	Description string    // Description of the item
	Price       float64   // Price of the item
	Stock       int       // Stock quantity of the item
	Status      string    // Current status of the item
	CreatedAt   time.Time // Time when the item was created
	UpdatedAt   time.Time // Time when the item was last updated
}

// ItemRepositoryPort defines the interface for item repository operations.
// It acts as a connector between the repository and the rest of the application.
type ItemRepositoryPort interface {
	SaveItem(*Item) (*Item, error)       // Save or update an item
	GetAllItems() (MapRepo, error)       // Retrieve all items
	GetItemByCode(string) (*Item, error) // Retrieve an item by its code
	GetItem(ID) (*Item, error)           // Retrieve an item by its ID
}
