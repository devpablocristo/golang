package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"restful_api_2/controllers"
	"restful_api_2/driver"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {

	/*
		no entiendo pq se armo aquí la conex con la db,
		cuando llamo desde una función da error

		pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
		logErrores(err)

		db, err := sql.Open("postgres", pgUrl)
		logErrores(err)
	*/

	db = driver.ConectarDb()
	controller := controllers.Controller{}

	r := mux.NewRouter()

	//no entiendo el uso de las funciones
	//entender las llamadas y las cargas
	r.HandleFunc("/", controller.Index())
	r.HandleFunc("/libros", controller.ObtenerLibros(db)).Methods("GET")
	r.HandleFunc("/libros/{id}", controller.ObtenerLibro(db)).Methods("GET")
	r.HandleFunc("/libros", controller.AñadirLibro(db)).Methods("POST")
	r.HandleFunc("/libros", controller.ActualizarLibro(db)).Methods("PUT")
	r.HandleFunc("/libros/{id}", controller.EliminarLibro(db)).Methods("DELETE")

	fmt.Println("El servidor corre en el puerto 8080")
	err := http.ListenAndServe(":8080", r)
	logErrores(err)
}

func logErrores(err error) {
	if err != nil {
		log.Fatal("ERROR!!!", err)
	}
}
