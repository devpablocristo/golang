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

	/*
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
			Title: "El 'Martín Fierro'",
			Price: 42.75,
			ISBN:  "84-206-1933-7",
		}
	*/

	Inventory = []InventoryInfo{
		{
			Title: "Fundation",
			Stock: 41,
		},
		{
			Title: "Solaris",
			Stock: 32,
		},
		{
			Title: "Rendezvous with Rama",
			Stock: 12,
		},
		{
			Title: "Neuromancer",
			Stock: 93,
		},
	}

	router := mux.NewRouter()

	const port string = ":8000"

	router.HandleFunc("/inventory", listBooks).Methods("GET")
	router.HandleFunc("/inventory", addBooks).Methods("POST")
	//router.HandleFunc("/inventory", editBook).Methods("PUT")
	//router.HandleFunc("/inventory", deleteBook).Methods("DELETE")

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

	// fmt.Fprintln(w, "Invetory:", Inventory)
	// fmt.Fprintln(w, "inventoryMap:", inventoryMap)
	// fmt.Fprintln(w, "---------------------------")

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
	//fmt.Fprintln(w, "diff:", diff)
	// fmt.Fprintln(w, reflect.TypeOf(Inventory))
	// fmt.Fprintln(w, reflect.TypeOf(newBooksSlice))
	// fmt.Fprintln(w, reflect.TypeOf(newBooksMap))
	// fmt.Fprintln(w, reflect.TypeOf(diff))
}
