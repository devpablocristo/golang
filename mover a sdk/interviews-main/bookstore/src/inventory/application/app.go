package application

import (
	"strings"

	inventory "github.com/devpablocristo/interviews/bookstore/src/inventory/domain"
)

type InventoryApp interface {
	GetBook(i inventory.InventoryInfo) inventory.InventoryInfo
}

func GetBook(i inventory.InventoryInfo) inventory.InventoryInfo {
	i.Book.Title = strings.ToLower(i.Book.Title)
	return i
}
