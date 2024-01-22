package main

import (
	"fmt"
	"time"
)

func main() {
	h, m, s := time.Now().Clock()
	fmt.Println("Hour: ", h)
	fmt.Println("Minute: ", m)
	fmt.Println("Second: ", s)
}
