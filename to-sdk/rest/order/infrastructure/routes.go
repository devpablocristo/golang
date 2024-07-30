


router.HandleFunc("/orders", listOrders).Methods("GET")
router.HandleFunc("/orders", createOrder).Methods("POST")