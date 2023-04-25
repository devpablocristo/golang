package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("test_write2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write String to file
	bytes2, err := file.WriteString("Writing string\n")
	if err != nil {
		panic(err)
	}

	// Write bytes of string to file
	bytes1, err := file.Write([]byte("Hello Golang dev\n"))
	if err != nil {
		panic(err)
	}

	file.Sync() // flush

	// Write using buffer
	buffer := bufio.NewWriter(file)
	bytes3, err := buffer.WriteString("Writting buffered\n")
	if err != nil {
		panic(err)
	}

	buffer.Flush() // Flush writes any buffered ..

	fmt.Printf("Wrote 1 %d bytes.\n", bytes1)
	fmt.Printf("Wrote 2 %d bytes.\n", bytes2)
	fmt.Printf("Wrote 3 %d bytes.\n", bytes3)
}
