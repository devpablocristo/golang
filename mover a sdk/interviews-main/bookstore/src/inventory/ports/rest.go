package port

import (
	application "github.com/devpablocristo/interviews/bookstore/src/inventory/application"
	inventory "github.com/devpablocristo/interviews/bookstore/src/inventory/domain"
)

type RestService struct {
	app application.InventoryApp
}

func (a RestService) GetBook() inventory.InventoryInfo {
	return a.app.GetBook(inventory.InventoryInfo{})
}
