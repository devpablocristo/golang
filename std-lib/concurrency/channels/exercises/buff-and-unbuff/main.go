package main

import (
	"fmt"
)

func main() {
	chUnbuff := make(chan int)
	chBuff := make(chan int, 3)
	//var wg sync.WaitGroup

	// chBuff <- 90
	// chBuff <- 90
	// chBuff <- 90
	// chBuff <- 90

	// T, bool chan
	// bool true abierto
	// booo false cerrado

	//wg.Add(1)
	go foo(chUnbuff)
	//go bar(&wg, chBuff)
	go bar(chBuff)

	//wg.Wait()

	for i := 0; i < 10; i++ {
		fmt.Println(<-chUnbuff)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-chBuff)
	}

	//fmt.Println("vacio?:", <-chBuff)
	bar2(chBuff)
	close(chBuff)

	_, ok := <-chBuff
	if !ok {
		for i := 0; i < 2; i++ {
			fmt.Println(<-chBuff)
		}
		fmt.Println("cerrado")
	} else {
		fmt.Println("abierto")
	}

}

func foo(chUnbuf chan int) {
	for i := 0; i < 10; i++ {
		chUnbuf <- i // cola fifo
	}

}

func bar( /*wg *sync.WaitGroup, */ chBuff chan int) {
	//defer wg.Done()
	for i := 10; i < 20; i++ {
		chBuff <- i
	}
}

func bar2(chBuff chan int) {
	for i := 20; i < 21; i++ {
		chBuff <- i
	}
	close(chBuff)
}
