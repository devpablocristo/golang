package main

import (
	"os"
	"sync"

	orderManager "github.com/devpablocristo/99minutos/order-manager/api"
)

const (
	defaultPort = "8080"
	port1       = "8081"
	port2       = "8082"
)

func main() {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	reserveNumberPort := os.Getenv("number-manager_PORT")
	if reserveNumberPort == "" {
		reserveNumberPort = defaultPort
	}

	wg.Add(1)
	go orderManager.StartApi(&wg, reserveNumberPort)
	wg.Wait()
}
