package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	repo "github.com/devpablocristo/golang-docker-tutorial/repo"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func UserPage(w http.ResponseWriter, r *http.Request) {
	users := repo.GetUsers()

	fmt.Println("Endpoint Hit: usersPage")
	json.NewEncoder(w).Encode(users)
}
