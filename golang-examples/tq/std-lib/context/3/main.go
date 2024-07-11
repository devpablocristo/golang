package main

import (
	"context"
	"fmt"
	"time"

	users "github.com/devpablocristo/golang/std-lib/context/3/users"
)

func main() {

	// main goroutine

	// universo a
	ctx := context.Background() // ctx vacio

	ctx, _ = context.WithTimeout(ctx, time.Second*10)

	usersFinishCan := make(chan bool)

	go users.ListUsers(ctx, usersFinishCan)

	select {
	case <-usersFinishCan:
		fmt.Println("Users finished")
	}
}
