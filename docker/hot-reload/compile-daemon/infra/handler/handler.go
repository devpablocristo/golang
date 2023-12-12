package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	repo "compile-daemon/infra/repo"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola!!!!!!!!!")
	fmt.Println("Endpoint Hit: homePage")
}

func UserPage(w http.ResponseWriter, r *http.Request) {
	users := repo.GetUsers()

	fmt.Println("Endpoint Hit: usersPage")
	json.NewEncoder(w).Encode(users)
}
