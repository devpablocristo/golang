package item

type ItemRepositoryPort interface {
	SaveItem(Item) error
	ListItems() (MapRepo, error)
}
