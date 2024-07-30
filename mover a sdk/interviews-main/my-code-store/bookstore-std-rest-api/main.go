package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

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

type Order struct {
	Customer Person       `json:"client"`
	Date     time.Time    `json:"date"`
	Details  []OrderItems `json:"details"`
}

type OrderItems struct {
	Book     Book  `json:"books"`
	Quantity int64 `json:"quantity"`
}

var Inventory []InventoryInfo
var Orders []Order

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
		Title: "El 'Martín Fierro'",
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

	customer1 := Person{
		Firstname: "Juan",
		Lastname:  "Perez",
	}

	customer2 := Person{
		Firstname: "John",
		Lastname:  "Smith",
	}

	details1 := []OrderItems{
		{
			Book:     book1,
			Quantity: 3,
		},
		{
			Book:     book2,
			Quantity: 1,
		},
	}

	details2 := []OrderItems{
		{
			Book:     book3,
			Quantity: 2,
		},
		{
			Book:     book4,
			Quantity: 4,
		},
	}

	Orders = []Order{
		{
			Customer: customer1,
			Date:     time.Now(),
			Details:  details1},
		{
			Customer: customer2,
			Date:     time.Now(),
			Details:  details2},
	}

	router := mux.NewRouter()

	const port string = ":8000"

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/inventory", listBooks).Methods("GET")
	router.HandleFunc("/inventory", addBooks).Methods("POST")
	router.HandleFunc("/orders", listOrders).Methods("GET")
	router.HandleFunc("/orders", createOrder).Methods("POST")

	log.Println("Server listining on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Home Page")
}

// List all books on the inventory
func listBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Inventory)
}

// Add books to inventory
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

// List all book orders
func listOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Orders)
}

// Creates a order of book
func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order

	err := json.NewDecoder(r.Body).Decode(&order)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error":"Error unmarshaling the request: %v"}`, err)
		return
	}

	Orders = append(Orders, order)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Orders)
}
