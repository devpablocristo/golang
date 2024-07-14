package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	ID     int      `json:"id "`
	Name   string   `json:"name"`
	Phones []string `json:"phones"`
}

func main() {

	// Simple string to JSON
	value := "golang"
	jsonValue, _ := json.Marshal(value)
	fmt.Println("Simple string to JSON:")
	fmt.Println(string(jsonValue))

	// Struct to JSON
	user := User{ID: 123456, Name: "Elvis", Phones: []string{"1234", "2222", "3333"}}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Struct to JSON:")
	fmt.Println(string(jsonUser))

	// JSON to Structs
	var userFromJson User
	userJson := `{"id ":123456,"name":"Elvis","phones":["1234","2222","3333"]}`
	err = json.Unmarshal([]byte(userJson), &userFromJson)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("JSON to Struct:")
	fmt.Println(userFromJson)

}
