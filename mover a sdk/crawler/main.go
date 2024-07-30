package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

//<a\s+(?:[^>]*?\s+)?href="([^"]*)">

func countDonwt(a int) int {
	var b int

	if a > 0 {
		fmt.Println(a)
		b := a - 1
		countDonwt(b)
	}

	return b
}

func main() {

	countDonwt(10)

	// s, err := GetHTML("https://example.com/")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(s)

	resultC, errC := recJob(ctx, workers, input) // returns results & `error` channels

	// asynchronous results - so read that channel first in the event of partial results ...
	for r := range resultC {
		fmt.Println(r)
	}

	// ... then check for any errors
	if err := <-errC; err != nil {
		log.Fatal(err)
	}

}

func GetHTML(getBodyURL string) (string, error) {
	resp, err := http.Get(getBodyURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	sbody := string(body)

	return sbody, nil

}

var (
	workers = 10
	ctx     = context.TODO() // use request context here - otherwise context.Background()
	input   = "abc"
)

func recJob(ctx context.Context, workers int, input string) (resultsC <-chan string, errC <-chan error) {

	// RW channels
	out := make(chan string)
	eC := make(chan error, 1)

	// R-only channels returned to caller
	resultsC, errC = out, eC

	// create workers + waitgroup logic
	go func() {

		var err error // error that will be returned to call via error channel

		defer func() {
			close(out)
			eC <- err
			close(eC)
		}()

		var wg sync.WaitGroup
		wg.Add(1)
		in := make(chan string) // input channel: shared by all workers (to read from and also to write to when they need to delegate)

		workerErrC := createWorkers(ctx, workers, in, out, &wg)

		// get the ball rolling, pass input job to one of the workers
		// Note: must be done *after* workers are created - otherwise deadlock
		in <- input

		errCount := 0

		// wait for all worker error codes to return
		for err2 := range workerErrC {
			if err2 != nil {
				log.Println("worker error:", err2)
				errCount++
			}
		}

		// all workers have completed
		if errCount > 0 {
			err = fmt.Errorf("PARTIAL RESULT: %d of %d workers encountered errors", errCount, workers)
			return
		}

		log.Printf("All %d workers have FINISHED\n", workers)
	}()

	return
}

func createWorkers(ctx context.Context, workers int, in chan string, out chan<- string, rwg *sync.WaitGroup) (errC <-chan error) {

	eC := make(chan error) // RW-version
	errC = eC              // RO-version (returned to caller)

	// track the completeness of the workers - so we know when to wrap up
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		i := i
		go func() {
			defer wg.Done()

			var err error

			// ensure the current worker's return code gets returned
			// via the common workers' error-channel
			defer func() {
				if err != nil {
					log.Printf("worker #%3d ERRORED: %s\n", i+1, err)
				} else {
					log.Printf("worker #%3d FINISHED.\n", i+1)
				}
				eC <- err
			}()

			log.Printf("worker #%3d STARTED successfully\n", i+1)

			// worker scans for input
			for input := range in {

				err = recFunc(ctx, input, in, out, rwg)
				if err != nil {
					log.Printf("worker #%3d recurseManagers ERROR: %s\n", i+1, err)
					return
				}
			}

		}()
	}

	go func() {
		rwg.Wait() // wait for all recursion to finish
		close(in)  // safe to close input channel as all workers are blocked (i.e. no new inputs)
		wg.Wait()  // now wait for all workers to return
		close(eC)  // finally, signal to caller we're truly done by closing workers' error-channel
	}()

	return
}
