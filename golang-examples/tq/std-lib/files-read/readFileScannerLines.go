package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("file_read_lines.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file) // default line by line
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
