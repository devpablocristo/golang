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

var Persons []Person

func main() {

	p1 := &Person{"Lucia", 26, []string{"Leer", "Pintar", "Ir al gym", "Estudiar inglés"}}
	p2 := &Person{"Maria", 22, []string{"Pintar", "Ir al gym", "Estudiar inglés"}}
	p3 := &Person{"Emilia", 24, []string{"Bailar", "Ir al gym", "Estudiar francés"}}

	Persons = append(Persons, *p1, *p2, *p3)

	loadTemplate("template2.txt", map[string]interface{}{"pers": Persons})

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
