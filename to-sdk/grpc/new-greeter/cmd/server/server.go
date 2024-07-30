package main

import (
	"context"
	"os"
	"sync"

	"time"

	"github.com/joho/godotenv"

	ctypes "github.com/devpablocristo/qh/internal/platform/custom-types"
)

func main() {
	rootDir, err := os.Getwd()
	if err != nil {
		ctypes.HandleFatalError("error getting current working directory: %v", err)
	}

	envFile := rootDir + "/.env"
	err = godotenv.Load(envFile)
	if err != nil {
		ctypes.HandleFatalError("error loading .env file", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Start(ctx)
}

func Start(ctx context.Context) {
	app := NewAppLauncher()
	app.Setup(ctx)
	var wg sync.WaitGroup

	// Start the REST service in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := app.InitEventService(ctx); err != nil {
			ctypes.HandleFatalError("error starting the REST service", err)
		}
	}()

	// Start the gRPC service in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := app.InitGreeterServer(ctx); err != nil {
			ctypes.HandleFatalError("error starting gRPC services", err)
		}
	}()

	wg.Wait()
}
