package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const port = ":9000"

type Book struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
	Price  int    `json:"price"`
}

var db []Book

func main() {
	b1 := Book{
		ID:     1,
		Title:  "Dune",
		Price:  1965,
		Author: "Frank Herbert",
	}

	b2 := Book{
		ID:     2,
		Title:  "Cita con Rama",
		Price:  1974,
		Author: "Arthur C. Clarke",
	}

	b3 := Book{
		ID:     3,
		Title:  "Un guijarro en el cielo",
		Price:  500,
		Author: "Isaac Asimov",
	}

	db = append(db, b1, b2, b3)

	r := gin.Default()

	r.GET("/", index)
	r.GET("/books", getBooks)
	r.POST("/books", addBook)
	r.GET("/books/:id", getBook)
	r.PUT("/books/:id", updateBook)
	r.DELETE("/books/:id", deleteBook)

	log.Println("Server listening on port", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalln(err)
	}
}

/*
w response: respuesta del servidor al cliente
r request: peticion del cliente al servidor
*/
func index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Bienvenido a mi increible API!")
}

func getBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  db,
	})
}

func getBook(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	for _, v := range db {
		if v.ID == id {
			ctx.JSON(http.StatusOK, gin.H{
				"error": false,
				"data":  v,
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": true,
		"data":  "book not found",
	})
}

func addBook(ctx *gin.Context) {
	request := ctx.Request

	var book Book
	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	db = append(db, book)

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  db,
	})
}

func updateBook(ctx *gin.Context) {
	r := ctx.Request
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	var book Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	for i, v := range db {
		if v.ID == id {
			db[i] = book
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  db,
	})
}

func deleteBook(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	for i, v := range db {
		if v.ID == id {
			db = append(db[:i], db[i+1:]...)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  db,
	})
}
