package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"restful_api_2/models"
	libroRepository "restful_api_2/repository/libro"
	"restful_api_2/utils"
	"strconv"

	"github.com/gorilla/mux"
)

type Controller struct{}

type libros []models.Libro

var sl libros

func (c Controller) ObtenerLibros(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r // para que no salga el warning, suponog que no hace falta
		var lib models.Libro

		sl = nil
		libroRepo := libroRepository.LibroRepository{}
		sl, err := libroRepo.ObtenerLibros(db, lib, sl)
		logErrores(err, w)

		w.Header().Set("Content-Type", "application/json")
		utils.EnviarExito(w, sl)
	}
}

func (c Controller) Index() http.HandlerFunc {
	/*
		w respuesta del servidor al cliente
		r peticion del cliente al servidor
	*/
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r // para que no salga el warning, suponog que no hace falta

		fmt.Fprintf(w, "Bienvenido a mi increible API!")
	}
}

func (c Controller) ObtenerLibro(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		_ = r // para que no salga el warning, suponog que no hace falta
		var lib models.Libro

		params := mux.Vars(r)

		sl = nil
		libroRepo := libroRepository.LibroRepository{}
		id, _ := strconv.Atoi(params["id"])

		lib, err := libroRepo.ObtenerLibro(db, lib, id)
		logErrores(err, w)

		w.Header().Set("Content-Type", "application/json")
		utils.EnviarExito(w, lib)
	}
}

func (c Controller) AñadirLibro(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*
			pegar en postman
			body -> ray -> paste!

			{
				"titulo": "Neuromante",
				"autor":  "William Gibson",
				"año":    "1985"
			}

			nota:
				· si termina con "/" no funca
				· solo si es asi "localhost:8080/libros"
		*/

		var lib models.Libro
		var libID int
		var error models.Error

		json.NewDecoder(r.Body).Decode(&lib)

		if lib.Autor == "" || lib.Titulo == "" || lib.Año == "" {
			error.Mensaje = "Enter missing fields."
			utils.EnviarError(w, http.StatusBadRequest, error) //400
			return
		}

		libroRepo := libroRepository.LibroRepository{}
		libID, err := libroRepo.AñadirLibro(db, lib)
		logErrores(err, w)

		w.Header().Set("Content-Type", "text/plain")
		utils.EnviarExito(w, libID)

	}
}

func (c Controller) ActualizarLibro(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*
			pegar en postman
			body -> ray -> paste!

			{
				"id": 1,
				"titulo": "Neuromante",
				"autor":  "William Gibson",
				"año":    "1985"
			}
		*/

		var lib models.Libro
		var error models.Error

		json.NewDecoder(r.Body).Decode(&lib)

		if lib.ID == 0 || lib.Autor == "" || lib.Titulo == "" || lib.Año == "" {
			error.Mensaje = "All fields are required."
			utils.EnviarError(w, http.StatusBadRequest, error)
			return
		}

		libroRepo := libroRepository.LibroRepository{}
		rowsUpdated, err := libroRepo.ActualizarLibro(db, lib)
		logErrores(err, w)

		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (c Controller) EliminarLibro(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var error models.Error
		params := mux.Vars(r)
		libroRepo := libroRepository.LibroRepository{}
		id, _ := strconv.Atoi(params["id"])

		rowsDeleted, err := libroRepo.EliminarLibro(db, id)

		if err != nil {
			error.Mensaje = "Server error."
			utils.EnviarError(w, http.StatusInternalServerError, error) //500
			return
		}

		if rowsDeleted == 0 {
			error.Mensaje = "Not Found"
			utils.EnviarError(w, http.StatusNotFound, error) //404
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.EnviarExito(w, rowsDeleted)
	}
}

func logErrores(err error, w http.ResponseWriter) {
	var error models.Error
	if err != nil {
		if err == sql.ErrNoRows {
			error.Mensaje = "Not found"
			utils.EnviarError(w, http.StatusNotFound, error)
			return
		} else {
			error.Mensaje = "Server error"
			utils.EnviarError(w, http.StatusInternalServerError, error)
			return
		}
	}
}
