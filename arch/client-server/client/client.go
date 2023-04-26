package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// TODO: connect to server on localhost port 8000
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	mustCopy(os.Stdout, conn)

}

// mustCopy - utility function
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
