package item

type ItemRepositoryPort interface {
	SaveItem(Item *Item) error
	ListItems() (MapRepo, error)
}
