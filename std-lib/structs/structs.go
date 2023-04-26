package main

import "fmt"

type User struct {
	Name     string
	LastName string
}

func main() {

	u := User{"Albert", "Einstein"}
	fmt.Println(u)

	fmt.Println(u.Name, "was a genius")
}
