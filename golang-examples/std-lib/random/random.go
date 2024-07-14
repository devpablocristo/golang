package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// change seed based on time
	rand.Seed(time.Now().UnixNano())

	// random int 0 <= n < 100
	fmt.Println("Random int number", rand.Intn(100))

	// random float64 0.0 <= f < 1.0
	fmt.Println("Random float64 number", rand.Float64())

	// random float32 0.0 <= f < 1.0.
	fmt.Println("Random float32 number", rand.Float32())

	// ..

	// change seed based on time
	rand.Seed(time.Now().UnixNano())

	// random int 0 <= n < 100
	fmt.Println("Random int number again", rand.Intn(100))

}
