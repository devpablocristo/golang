package main

import (
	"fmt"

	user "exer/user"
)

func main() {

	u := user.User{
		Name:     "Albert",
		LastName: "Einstein",
		Age:      76,
		UserAddress: user.Address{
			Street: "1234 Science Blvd",
			number: 42, // no accesible
		},
	}

	fmt.Println(u)

	fmt.Println(u.Name, "was a genius")
}
