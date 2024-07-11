package main

import (
	"fmt"
	"time"
)

type Job struct {
	id     int
	value  int
	value2 int
}
type Result struct {
	id        int
	sumValues int
}

func worker(id int, jobs <-chan Job, results chan<- Result) {
	for j := range jobs {
		fmt.Println("Job id", id, "started  job", j)
		time.Sleep(time.Second) // simulate long process
		result := Result{id: j.id, sumValues: j.value + j.value2}
		fmt.Println("Job id", id, "finished job", result)
		results <- result
	}
}

func main() {
	// job to send to worker
	jobs := make(chan Job, 100)

	// result from worker
	results := make(chan Result, 100)

	goroutines := 5
	// Create 5 workers
	for i := 1; i <= goroutines; i++ {
		go worker(i, jobs, results)
	}

	// Send 5 jobs using the channel jobs
	for j := 1; j <= goroutines; j++ {
		j := Job{id: j, value: j, value2: j}
		jobs <- j
	}
	close(jobs)

	// Receive the result from work
	for x := 1; x <= goroutines; x++ {
		<-results
	}
}
