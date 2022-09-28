package main

import (
	"os"
	"text/template"
)

type Person struct {
	Nombre  string
	Edad    int32
	Hobbies []string
}

var funcs = template.FuncMap{
	"increment": func(num int) int {
		return num + 1
	},
}

func main() {

	p1 := &Person{"Lucia", 26, []string{"Leer", "Pintar", "Ir al gym", "Estudiar inglés"}}

	loadTemplate("template2.txt", p1)

}

func loadTemplate(fileName string, data interface{}) {
	t, err := template.New(fileName).Funcs(funcs).ParseFiles("templates/" + fileName)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
