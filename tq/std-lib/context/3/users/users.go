package users

import (
	"context"
	"fmt"
	"time"
)

type User struct {
	Id       int
	Username string
}

var users []User = []User{
	{Id: 1, Username: "Mau"},
	{Id: 2, Username: "Gonza"},
	{Id: 3, Username: "Lu"},
	{Id: 4, Username: "Pablo"},
}

func ListUsers(ctx context.Context, ch chan<- bool) {
	for _, user := range users {
		contextIsDone := false

		select {
		case <-ctx.Done():
			contextIsDone = true
			break
		case <-time.After(time.Second):
			fmt.Println(user)
		}

		if contextIsDone {
			fmt.Println("context done")
			//ch <- true
			break
		}
	}

	ch <- true
}
