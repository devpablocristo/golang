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

	err := tpl.Execute(os.Stdout, lennon)
	if err != nil {
		log.Fatalln(err)
	}

}
