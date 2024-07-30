package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetInput() (string, error) {

	fmt.Println("\nCommands:")
	fmt.Println("-p number of parallel workers to navigate through site.")
	fmt.Println("-of output file path.")
	fmt.Println("-md  max depth of url navigation recursion.")
	fmt.Println("-e  exit program.")
	fmt.Println("")
	fmt.Print("sitemap-> ")
	s, err := reader.ReadString('\n')
	s = strings.Replace(s, "\n", "", 1)

	if err != nil {
		log.Fatal(err)
	}
	s = strings.Replace(s, "\n", "", 1)

	return s, nil
}

func ParseCmd(s string) (string, string) {
	var l, n []rune
	for _, r := range s {
		switch {
		case r >= 'A' && r <= 'Z':
			l = append(l, r)
		case r >= 'a' && r <= 'z':
			l = append(l, r)
		case r >= '0' && r <= '9':
			n = append(n, r)
		}
	}
	if string(l) == "e" {
		return "-e", ""
	}
	return "-" + string(l) + "=", string(n)
}
