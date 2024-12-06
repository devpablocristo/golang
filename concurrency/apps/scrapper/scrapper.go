package main

import (
	"fmt"
	"math/rand"
	"time"
)

type position struct {
	x int
	y int
	z int
}

type drone struct {
	UUID int
	battery  int
	position position
}

type droneStatus struct {
	battery  int
	distance int
}

 

func newDrone(id int) drone {
	return drone{
		UUID: id,
		battery: rand.Intn(100),
		position: position{
			x: 0,
			y: 0,
			z: 0,
		},
	}
}

func doDroneThings(dr drone, droneCom chan droneStatus) {
	for {
		select
			
	}
}

func work(worker chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		jobs <- i
	}
	fmt.Println("Cerre un canal")
	close(jobs)
}

func observer(dron chan int, done chan bool) {
	for number := range jobs {
		fmt.Println("Sacamos del jobs: ", number)
	}
	done <- true
}

func skynet() {
	droneCom := make(chan droneStatus, 1)

	dronesQuantity := 5
	for i := 0; i < dronesQuantity; i++ {
		go doDroneThings(newDrone(i), droneCom)
	}

	go assignJobs(jobs)

	done := make(chan bool, 1)

	go emptyChannel(jobs, done)

	<-done

	fmt.Println(<-jobs)

}
