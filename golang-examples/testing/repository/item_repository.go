package repository

//go:generate mockgen -source=./item_repository.go -destination=../mocks/item_repository_mock.go -package=mocks
type ItemRepository interface {
	SaveItem(name string, stock int) error
	GetItemByID(itemID uint) error
}
