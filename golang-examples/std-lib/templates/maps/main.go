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

	countries := map[string]string{
		"Argentina": "America",
		"India":     "Asia",
		"Nigeria":   "Africa",
		"Italia":    "Europe",
	}
	err := tpl.Execute(os.Stdout, countries)
	if err != nil {
		log.Fatalln(err)
	}

}
