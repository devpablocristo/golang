// Package main is the main entry point for the Go program.
package main

// Import necessary packages.
import (
	"log"      // Standard Go logging package.
	"net/http" // HTTP package for handling web requests.

	"github.com/gin-gonic/gin" // Gin is a web framework for Go.
)

// The main function, where the program execution begins.
func main() {

	// Create a new default Gin router.
	router := gin.Default()

	// Define a route for the root path ("/") that calls the helloWorld function.
	router.GET("/", helloWorld)

	// Log a message indicating that the server has started.
	log.Println("Server started at http://localhost:8080/")

	// Run the server on port 8080.
	err := router.Run(":8080")
	if err != nil {
		// Log a fatal message if there's an error starting the server.
		log.Fatal(err)
	}
}

// The helloWorld function, which will be called when a request is made to the root path ("/").
func helloWorld(c *gin.Context) {
	// Respond with a simple "Hello World!" message and a status code of 200 OK.
	c.String(http.StatusOK, "Hello, World!")
}
