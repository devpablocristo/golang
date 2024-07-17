package item

// ItemRepositoryPort define la interfaz para el repositorio de elementos
type ItemRepositoryPort interface {
	SaveItem(*Item) error
	ListItems() (MapRepo, error)
}
