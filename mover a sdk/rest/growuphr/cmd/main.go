package main

import (
	"os"
	"sync"

	numberManager "github.com/devpablocristo/growuphr/number-manager/api"
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

	wg.Add(3)
	go numberManager.StartApi(&wg, reserveNumberPort)
	go numberManager.StartApi(&wg, port1)
	go numberManager.StartApi(&wg, port2)

	wg.Wait()
}
