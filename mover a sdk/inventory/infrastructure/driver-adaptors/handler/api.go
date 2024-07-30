package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	domain "github.com/devpablocristo/golang/06-apps/qh/inventory/domain"
	mapdb "github.com/devpablocristo/golang/06-apps/qh/inventory/infrastructure/repository/inmemory/mapdb"
	slicedb "github.com/devpablocristo/golang/06-apps/qh/inventory/infrastructure/repository/inmemory/slicedb"
	application "github.com/devpablocristo/golang/06-apps/qh/person/application"
	pdomain "github.com/devpablocristo/golang/06-apps/qh/person/domain"
)

var (
	mapDB   = mapdb.NewMapDB()
	sliceDB = slicedb.NewSliceDB()
)

func init() {
	book1 := domain.Book{
		Author: pdomain.Person{
			Firstname: "Isaac",
			Lastname:  "Asimov",
		},
		Title: "Fundation",
		Price: 28.50,
		ISBN:  "0-553-29335-4",
	}

	book2 := domain.Book{
		Author: pdomain.Person{
			Firstname: "Stanislaw",
			Lastname:  "Lem",
		},
		Title: "Solaris",
		Price: 65.20,
		ISBN:  "0156027607",
	}

	book3 := domain.Book{
		Author: pdomain.Person{
			Firstname: "Arthur C.",
			Lastname:  "Clarck",
		},
		Title: "Rendezvous with Rama",
		Price: 53.50,
		ISBN:  "0-575-01587-X",
	}

	book4 := domain.Book{
		Author: pdomain.Person{
			Firstname: "Jorge Luis",
			Lastname:  "Borges",
		},
		Title: "El Aleph",
		Price: 42.75,
		ISBN:  "84-206-1933-7",
	}

	domain.Inventory = []domain.BookStock{
		{
			Book:  book1,
			Stock: 41,
		},
		{
			Book:  book2,
			Stock: 32,
		},
		{
			Book:  book3,
			Stock: 12,
		},
		{
			Book:  book4,
			Stock: 93,
		},
	}

	mapDB.SaveBook(&b1)
	mapDB.SaveBook(&b2)

	sliceDB.SaveBook(b1)
	sliceDB.SaveBook(b2)

	fmt.Println(*mapDB)
	fmt.Println(sliceDB)
}

func Books(wg *sync.WaitGroup) {
	people := []domain.Person{
		p1,
		p2,
	}

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}

func StartApi(wg *sync.WaitGroup, port string) {
	defer wg.Done()

	// db := postgres.ConnectToDB()
	// defer db.Close()

	ps := application.NewPersonaApplication()
	handler := NewChiHandler(ps, port)

	handler.SetupChiRoutes()
	handler.RunChiServer()
}
