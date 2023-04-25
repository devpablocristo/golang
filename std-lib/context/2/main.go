package main

import (
	"context"
	"fmt"
	"time"
)

func enrichCtx(ctx context.Context) context.Context {
	return context.WithValue(ctx, "user", "joe")
}

func doStuff(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		default:
			fmt.Println("doing stuff")

		}
		time.Sleep(500 * time.Millisecond)
	}
}

type DemoStruct struct {
	Val int
}

func demo_func(s *DemoStruct) {
	s.Val = 1
}

func main() {

	ctx := context.Background() // start with empty context
	//ctx2 := context.TODO()      // start with context that is marked as "to-do", meaning that it is not yet complete

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second) // create a context with a timeout of 10 seconds
	defer cancel()

	ctx = enrichCtx(ctx)
	go doStuff(ctx) // cancel the context when the function exits

	select {
	case <-ctx.Done(): // if the context is done, then the timeout has been reached
		fmt.Println("final timeout")
		fmt.Println(ctx.Err())
	}

	time.Sleep(2 * time.Second) // wait 5 seconds

}
