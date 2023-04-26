package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	txt, err := ioutil.ReadFile("file_read.txt")
	if err != nil {
		log.Fatal(err)
	}

	str := string(txt) // convert to a 'string'
	fmt.Println(str)   // print the content
}
