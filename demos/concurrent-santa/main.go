package main

import (
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	numOfElf       = 12 // numero de elfos totales
	totalReindeers = 9  // renos totales
)

var (
	counter      uint32            // contador de elfos que tienen problemas
	santaChUp    = make(chan bool) // canal para Santa despierto
	santaChSleep = make(chan bool) // canal para Santa dormido
	s            = santa{
		mutex: sync.Mutex{},
	}
)

type santa struct {
	mutex sync.Mutex
}

func elfsAreWorking(extWG *sync.WaitGroup) {
	var wg sync.WaitGroup

	for i := 1; i <= numOfElf; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			log.Println("elf", i, "crafting a toy")
			if rand.Intn(3) == 0 {
				atomic.AddUint32(&counter, 1)
				go santaHelp(i)
			} else {
				log.Println("toy of elf", i, "is done")
			}
			time.Sleep(2 * time.Second)
		}(i)
	}
	wg.Wait()
	extWG.Done()
}

func santaHelp(i int) {
	log.Println("The toy of elf", i, "is broken, he needs help")
	if atomic.LoadUint32(&counter) == 3 {
		santaChUp <- true
		time.Sleep(1 * time.Second)
		log.Println("Santa fixing the toys")
		helptime := time.Duration(2+rand.Intn(3)) * time.Second
		time.Sleep(helptime)
		log.Println("Problem fixed in", helptime, "seconds")
		santaChSleep <- true
		s.mutex.Unlock()
		atomic.StoreUint32(&counter, 0)
	}
}

func reindeerArrival(wg *sync.WaitGroup) {
	for i := 1; i <= totalReindeers; i++ {
		time.Sleep(time.Duration(5+rand.Intn(2)) * time.Second)
		log.Println("reindeer", i, "arrived!")
	}
	log.Println("All the reindeers are here")
	wg.Done()
}

func santaStateZero() {
	go func() {
		santaChSleep <- true
	}()
}

func santaRoutine() {
	for {
		select {
		case <-santaChUp:
			s.mutex.Lock()
			log.Println("Santa is up now...")
		case <-santaChSleep:
			log.Println("Santa is sleeping...")
		}
	}
}

func wakeUpSanta() {
	go func() { santaChUp <- true }()
}

func main() {
	var wg sync.WaitGroup

	santaStateZero()

	wg.Add(3)
	go santaRoutine()
	go elfsAreWorking(&wg)
	go reindeerArrival(&wg)

	wg.Wait()
	log.Println("Santa is ready to go!")
}
