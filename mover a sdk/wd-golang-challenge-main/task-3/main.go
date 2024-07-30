package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type InventoryInfo struct {
	Title string `json:"title"`
	Stock int64  `json:"stock"`
}

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
			Firstname: "William",
			Lastname:  "Gibson",
		},
		Title: "Neuromancer",
		Price: 42.75,
		ISBN:  "0-441-56956-0",
	}

	Inventory = []InventoryInfo{
		InventoryInfo{
			Book:  book1,
			Stock: 1,
		},
		InventoryInfo{
			Book:  book2,
			Stock: 2,
		},
		InventoryInfo{
			Book:  book3,
			Stock: 3,
		},
		InventoryInfo{
			Book:  book4,
			Stock: 4,
		},
	}

	router := mux.NewRouter()

	const port string = ":8000"

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
		inventoryMap[Inventory[i].Title] = Inventory[i].Stock
	}

	for _, book := range newBooksSlice {
		if _, found := inventoryMap[book.Title]; !found {
			Inventory = append(Inventory, book)
		} else {
			for i := range Inventory {
				if Inventory[i].Title == book.Title {
					Inventory[i].Stock += book.Stock
					break
				}
			}
		}
	}

	fmt.Fprintln(w, "Inventory:", Inventory)
}
