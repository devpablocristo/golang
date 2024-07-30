package handler

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"
)

func (h *Handler) StartServer() error {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	serverErr := make(chan error)

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := h.server.ListenAndServe(); err != nil {
			log.Println(err)
			serverErr <- err
		}
	}()
	close(serverErr)

	select {
	case err := <-serverErr:
		return err
	}

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	h.server.Shutdown(ctx)
	// Optionally, you could run h.server.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)

	return nil
}
