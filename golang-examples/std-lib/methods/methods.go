package main

import "fmt"

type User struct {
	Name     string
	LastName string
}

// this is a method
func (u User) CheckGenius() bool {
	return u.Name == "Albert"
}

// this is a common function
func CheckGenius(u User) bool {
	return u.Name == "Albert"
}

func main() {

	u := User{"Albert", "Einstein"}

	// using method throw User.CheckGenius
	fmt.Println("Method CheckGenius", u.Name, u.CheckGenius())

	// using a common function
	fmt.Println("Function CheckGenius", CheckGenius(u))

}
