package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

func HttpServer(port string, router http.Handler) {
	log.Println("starting chi server")

	httpServer := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	var wg sync.WaitGroup
	defer wg.Wait()

	wg.Add(1)
	go func() {
		err := httpServer.ListenAndServe()
		if err != nil {
			if err != http.ErrServerClosed {
				log.Fatalf("could not listen on %s due to %s", httpServer.Addr, err.Error())
			}
		}
		wg.Done()
	}()

	log.Printf("the chi server is ready to handle requests %s", httpServer.Addr)

}
