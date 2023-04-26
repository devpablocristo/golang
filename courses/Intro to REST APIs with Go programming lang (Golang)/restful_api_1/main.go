package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type libro struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
	Año    string `json:"año"`
}

type libros []libro

var sl libros

func main() {

	l1 := libro{
		ID:     1,
		Titulo: "Dune",
		Año:    "1965",
		Autor:  "Frank Herbert",
	}

	l2 := libro{
		ID:     2,
		Titulo: "Cita con Rama",
		Año:    "1974",
		Autor:  "Arthur C. Clarke",
	}

	l3 := libro{
		ID:     3,
		Titulo: "Un guijarro en el cielo",
		Año:    "1950",
		Autor:  "Isaac Asimov",
	}

	l4 := libro{
		ID:     4,
		Titulo: "Redshirts",
		Año:    "2013",
		Autor:  "John Scalzi",
	}

	l5 := libro{
		ID:     5,
		Titulo: "Hyperion",
		Año:    "1990",
		Autor:  "Dan Simmons",
	}

	//fmt.Println(l1)

	sl = append(sl, l1, l2, l3, l4, l5)

	//fmt.Println(sl)

	r := mux.NewRouter()

	//no entiendo el uso de las funciones
	//entender las llamadas y las cargas
	r.HandleFunc("/", index)
	r.HandleFunc("/libros", obtenerLibros).Methods("GET")
	r.HandleFunc("/libros/{id}", obtenerLibro).Methods("GET")
	r.HandleFunc("/libros", añadirLibro).Methods("POST")
	r.HandleFunc("/libros", actualizarLibro).Methods("PUT")
	r.HandleFunc("/libros/{id}", eliminarLibro).Methods("DELETE")

	err := http.ListenAndServe(":8080", r)

	//fmt.Println(err)
	//leer como se hacer bien el control de errores
	log.Fatal(err)
}

/*
w respuesta del servidor al cliente
r peticion del cliente al servidor
*/
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a mi increible API!")
}

func obtenerLibros(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(sl)
}

func obtenerLibro(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("obtener un libro")
	params := mux.Vars(r)
	// estudiar diferencias
	//fmt.Println(params)
	//log.Println(params)

	idRequest, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Error:", err)
	}

	for _, v := range sl {
		if v.ID == idRequest {
			json.NewEncoder(w).Encode(&v)
		}
		fmt.Println("la")
	}
}

func añadirLibro(w http.ResponseWriter, r *http.Request) {
	/*
		pegar en postman
		body -> ray -> paste!

		{
			"id":     6,
			"titulo": "Neuromante",
			"autor":  "William Gibson",
			"año":    "1985"
		}
	*/

	var l libro
	json.NewDecoder(r.Body).Decode(&l)

	//sl = append

	sl = append(sl, l)

	json.NewEncoder(w).Encode(sl)
}

func actualizarLibro(w http.ResponseWriter, r *http.Request) {
	var l libro

	// se carga l con la nueva info
	json.NewDecoder(r.Body).Decode(&l)

	for i, v := range sl {
		if v.ID == l.ID {
			sl[i] = l
		}
	}

	json.NewEncoder(w).Encode(sl)
}

func eliminarLibro(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	idRequest, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Error:", err)
	}

	for i, v := range sl {
		if v.ID == idRequest {
			sl = append(sl[:i], sl[i+1:]...)
		}
	}

	json.NewEncoder(w).Encode(sl)

}
