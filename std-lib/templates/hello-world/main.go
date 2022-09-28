package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// template.ParseFiles, can get 0 or more files

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// create index.html
	nf, err := os.Create("index.html")
	if err != nil {
		log.Println(err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
