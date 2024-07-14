package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

type Album struct {
	Title string
	Year  int
}

type Band struct {
	Members []Person
	Albums  []Album
}

func main() {
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))

	lennon := Person{
		Name: "John",
		Age:  30,
	}

	mcartney := Person{
		Name: "Paul",
		Age:  29,
	}

	star := Person{
		Name: "Ringo",
		Age:  25,
	}

	harrison := Person{
		Name: "George",
		Age:  31,
	}

	ps := []Person{lennon, mcartney, star, harrison}

	err := tpl.Execute(os.Stdout, ps)
	if err != nil {
		log.Fatalln(err)
	}
}
