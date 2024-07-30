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

func main() {
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))

	catNames := []string{"Toribio", "Nana", "Fortunata", "Oki"}
	err := tpl.Execute(os.Stdout, catNames)
	if err != nil {
		log.Fatalln(err)
	}

}
