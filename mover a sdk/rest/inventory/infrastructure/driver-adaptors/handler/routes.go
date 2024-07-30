package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	muxrouter "github.com/devpablocristo/golang/06-apps/bookstore/inventory/infrastructure/router"
)

var httpMuxRouter muxrouter.MuxRouter = *muxrouter.NewMuxRouter()

func (h Handler) SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/inventory/{id}", h.GetBookByISBN).Methods("GET")

	// router.HandleFunc("/inventory", listBooks).Methods("GET")
	// router.HandleFunc("/inventory", addBook).Methods("POST")
	// router.HandleFunc("/inventory/{id}", updateBook).Methods("PUT")
	// router.HandleFunc("/inventory/{id}", deleteBookByID).Methods("DELETE")
	// router.HandleFunc("/inventory/{id}", updateBookByPatch).Methods("PATCH")
	// router.HandleFunc("/inventoryISBN/{isbn}", isbn_containes).Methods("GET")

	// httpMuxRouter.post("/inventory/add", inventoryControllers.Add)
	// httpMuxRouter.get("/inventory/all", inventoryControllers.GetAll)

	const port string = ":8888"
	log.Println("Server listining on port", port)
	log.Fatalln(http.ListenAndServe(port, router))

	return router
}
