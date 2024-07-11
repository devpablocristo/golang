package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "Estatua de Messi")
	if err != nil {
		log.Fatalln(err)
	}
}
