package main

import (
	"io/ioutil"
)

func main() {

	err := ioutil.WriteFile("test_write.txt", []byte("Hello go"), 0666)
	if err != nil {
		panic(err)
	}
}
