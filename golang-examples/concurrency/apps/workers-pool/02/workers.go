package main

import (
	"fmt"
	"strconv"
	"time"
)

// Recibes a int, as id, and 2 channels
// jobs is a only recibe channel
// results is a only send channel
func worker(id int, jobs <-chan int, results chan<- int, msg string) {

	fmt.Println(msg)
	// the for operator will end when the channel is closed.
	// jobs recibe 5 values (ints) and it it exits the loop, because the channel jobs is closed.
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// jobs channels
	const numJobs = 5

	/*

		https://github.com/uber-go/guide/blob/master/style.md#channel-size-is-one-or-none

		Channels should usually have a size of one or be unbuffered. By default, channels are unbuffered and have a size of zero. Any other size must be subject to a high level of scrutiny. Consider how the size is determined, what prevents the channel from filling up under load and blocking writers, and what happens when this occurs.

		Bad	Good
		// Ought to be enough for anybody!
		c := make(chan int, 64)

		// Size of one
		c := make(chan int, 1) // or
		// Unbuffered channel, size of zero
		c := make(chan int)

	*/
	// create buffered chan jobs
	jobs := make(chan int, 1)

	// create buffered chan results
	results := make(chan int, 1)

	// creates 3 goroutnes
	// sends 3 ints and 2 chans (jobs and result)
	msg := ""
	for w := 1; w <= 3; w++ {
		msg = strconv.Itoa(w) + " sent worker"
		// each go routine recibes 2 chans and and 3 diff workers idsone diff working
		go worker(w, jobs, results, msg)
	}

	// Sends 5 ints throw the channel jobs
	// jobs receive work on the jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// Close jobs
	/*
		It's OK to leave a Go channel open forever and never close it. When the channel is no longer used, it will be garbage collected.

		GC:  garbage collector is a piece of software that performs automatic memory management. Its job is to free any unused memory and ensure that no memory is freed while it is still in use

	*/
	close(jobs)

	// Recibes 5 ints throw the channel result
	// results send the corresponding results on results.
	for a := 1; a <= numJobs; a++ {
		fmt.Println("Results:", <-results)
	}
}
