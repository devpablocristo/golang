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

	hardDay := Album{
		Title: "A Hard Day's Night",
		Year:  1964,
	}

	help := Album{
		Title: "Help!",
		Year:  1965,
	}

	letItBe := Album{
		Title: "Let It Be",
		Year:  1970,
	}

	beatles := []Person{lennon, mcartney, star, harrison}

	albumsBeatles := []Album{hardDay, help, letItBe}

	bands := Band{
		Members: beatles,
		Albums:  albumsBeatles,
	}

	err := tpl.Execute(os.Stdout, bands)
	if err != nil {
		log.Fatalln(err)
	}
}
