package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, stop := context.WithCancel(context.Background())
	go func() {
		fmt.Scanln()
		stop()
	}()

	heartbeat(ctx)
}

func heartbeat(ctx context.Context) {
	tick := time.Tick(time.Second)

	for {
		select {
		case <-tick:
		case <-ctx.Done():
			return
		}
		fmt.Println("beat")
	}
}
