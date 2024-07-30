package adapter

import (
	"net/http"

	domain "github.com/devpablocristo/interviews/bookstore/src/inventory/domain"
)

// func ListBooks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(myRest.Inventory)
// }

func GetBook(w http.ResponseWriter, r *http.Request) domain.Book {

	book1 := domain.Book{
		Author: domain.Person{
			Firstname: "Isaac",
			Lastname:  "Asimov",
		},
		Title: "Fundation",
		Price: 28.50,
		ISBN:  "0-553-29335-4",
	}

	return book1

}

// func AddBooks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	received_JSON, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintln(w, err)
// 		return
// 	}

// 	newBooksSlice := make([]domain.InventoryInfo, 0)
// 	err = json.Unmarshal(received_JSON, &newBooksSlice)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintf(w, `{"error":"Error unmarshaling the request: %v"}`, err)
// 		return
// 	}

// 	inventoryMap := make(map[string]int64)
// 	for i := 0; i < len(domain.Inventory); i++ {
// 		inventoryMap[domain.Inventory[i].Book.ISBN] = domain.Inventory[i].Stock
// 	}

// 	for _, v := range newBooksSlice {
// 		if _, found := inventoryMap[v.Book.ISBN]; !found {
// 			domain.Inventory = append(domain.Inventory, v)
// 		} else {
// 			for i := range domain.Inventory {
// 				if domain.Inventory[i].Book.ISBN == v.Book.ISBN {
// 					domain.Inventory[i].Stock += v.Stock
// 					break
// 				}
// 			}
// 		}
// 	}

// 	// fmt.Fprintf(w, "inventoryMap: %+v", inventoryMap)
// 	// fmt.Fprintf(w, "Inventory: %+v", Inventory)

// 	fmt.Fprintln(w, "Inventory:")
// 	for _, v := range domain.Inventory {
// 		fmt.Fprintln(w, "Book:", v.Book.Title)
// 		fmt.Fprintln(w, "Stork:", v.Stock)
// 		fmt.Fprintln(w, "------------------")
// 	}

// }
