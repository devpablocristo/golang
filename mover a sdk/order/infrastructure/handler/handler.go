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
		Details:  details1
	},
	{
		Customer: customer2,
		Date:     time.Now(),
		Details:  details2 
	},
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
