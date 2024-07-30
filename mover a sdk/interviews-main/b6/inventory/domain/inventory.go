package inventory

type InventoryRepository interface {
	SaveBook(book Book) error
	ListInventory() ([]Book, error)
}

type Book struct {
	Author Person  `json:"author"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
	ISBN   string  `json:"isbn"`
}

type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type BookStock struct {
	Book  Book `json:"book"`
	Stock int  `json:"stock"`
}
