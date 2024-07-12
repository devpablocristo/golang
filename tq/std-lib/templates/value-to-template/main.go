package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "Mi gato se llama Toribio")
	if err != nil {
		log.Fatalln(err)
	}
}
