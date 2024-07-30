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

	Inventory = []InventoryInfo{
		{
			Title: "Fundation",
			Stock: 1,
		},
		{
			Title: "Solaris",
			Stock: 2,
		},
		{
			Title: "Rendezvous with Rama",
			Stock: 3,
		},
		{
			Title: "Neuromancer",
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
