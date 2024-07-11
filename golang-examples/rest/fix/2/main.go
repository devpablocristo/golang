package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Book struct {
	Author Person  `json:"author"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
	ISBN   string  `json:"isbn"`
}

type InventoryInfo struct {
	Book  Book  `json:"book"`
	Stock int64 `json:"stock"`
}

// type InventoryInfo struct {
// 	Title string `json:"title"`
// 	Stock int64  `json:"stock"`
// }

var Inventory []InventoryInfo

func main() {

	book1 := Book{
		Author: Person{
			Firstname: "Isaac",
			Lastname:  "Asimov",
		},
		Title: "Fundation",
		Price: 28.50,
		ISBN:  "0-553-29335-4",
	}

	book2 := Book{
		Author: Person{
			Firstname: "Stanislaw",
			Lastname:  "Lem",
		},
		Title: "Solaris",
		Price: 65.20,
		ISBN:  "0156027607",
	}

	book3 := Book{
		Author: Person{
			Firstname: "Arthur C.",
			Lastname:  "Clarck",
		},
		Title: "Rendezvous with Rama",
		Price: 53.50,
		ISBN:  "0-575-01587-X",
	}

	book4 := Book{
		Author: Person{
			Firstname: "Jorge Luis",
			Lastname:  "Borges",
		},
		Title: "El Aleph",
		Price: 42.75,
		ISBN:  "84-206-1933-7",
	}

	Inventory = []InventoryInfo{
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

	router := mux.NewRouter()

	const port string = ":8888"

	router.HandleFunc("/inventory", listBooks).Methods("GET")
	router.HandleFunc("/inventory", addBooks).Methods("POST")

	log.Println("Server listining on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}

func listBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Inventory)
}

func addBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	received_JSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	newBooksSlice := make([]InventoryInfo, 0)
	err = json.Unmarshal(received_JSON, &newBooksSlice)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error":"Error unmarshaling the request: %v"}`, err)
		return
	}

	inventoryMap := make(map[string]int64)
	for i := 0; i < len(Inventory); i++ {
		inventoryMap[Inventory[i].Book.ISBN] = Inventory[i].Stock
	}

	for _, v := range newBooksSlice {
		if _, found := inventoryMap[v.Book.ISBN]; !found {
			Inventory = append(Inventory, v)
		} else {
			for i := range Inventory {
				if Inventory[i].Book.ISBN == v.Book.ISBN {
					Inventory[i].Stock += v.Stock
					break
				}
			}
		}
	}

	// fmt.Fprintf(w, "inventoryMap: %+v", inventoryMap)
	// fmt.Fprintf(w, "Inventory: %+v", Inventory)

	fmt.Fprintln(w, "Inventory:")
	for _, v := range Inventory {
		fmt.Fprintln(w, "Book:", v.Book.Title)
		fmt.Fprintln(w, "Stork:", v.Stock)
		fmt.Fprintln(w, "------------------")
	}

}
