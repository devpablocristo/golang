package main

import (
	"os"
	"sync"

	cmsenv "github.com/devpablocristo/blankfactor/common/env"
	eventService "github.com/devpablocristo/blankfactor/event-service/api"
)

const (
	defaultPort = "8080"
)

func main() {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	cmsenv.LoadEnv()

	reserveNumberPort := os.Getenv("SERVER_PORT_1")
	if reserveNumberPort == "" {
		reserveNumberPort = os.Getenv(defaultPort)

	}

	wg.Add(1)
	go eventService.StartApi(&wg, reserveNumberPort)
	wg.Wait()
}
