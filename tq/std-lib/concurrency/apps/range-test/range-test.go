package main

import (
	"fmt"
	"time"
)

func assignJobs(jobs chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		jobs <- i
	}
	fmt.Println("Cerre un canal")
	close(jobs)
}

func emptyChannel(jobs chan int, done chan bool) {
	for number := range jobs {
		fmt.Println("Sacamos del jobs: ", number)
	}
	done <- true
}

func main() {
	jobs := make(chan int, 1)

	go assignJobs(jobs)

	done := make(chan bool, 1)

	go emptyChannel(jobs, done)

	<-done

	fmt.Println(<-jobs)

}
