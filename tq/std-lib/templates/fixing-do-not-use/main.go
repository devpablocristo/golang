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

	// jagger := Person{
	// 	Name: "Mick",
	// 	Age:  30,
	// }

	// richards := Person{
	// 	Name: "Keith",
	// 	Age:  29,
	// }

	// wood := Person{
	// 	Name: "Ronnie",
	// 	Age:  25,
	// }

	// watts := Person{
	// 	Name: "Charlie",
	// 	Age:  31,
	// }

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

	// letItBleed := Album{
	// 	Title: "Let It Bleed",
	// 	Year:  1969,
	// }

	// sticky := Album{
	// 	Title: "Sticky Fingers",
	// 	Year:  1971,
	// }

	// rockAndRoll := Album{
	// 	Title: "It's Only Rock 'n' Roll",
	// 	Year:  1974,
	// }

	beatles := []Person{lennon, mcartney, star, harrison}
	//stones := []Person{jagger, richards, wood, watts}

	albumsBeatles := []Album{hardDay, help, letItBe}
	//albumsStones := []Album{letItBleed, sticky, rockAndRoll}

	bands := Band{
		Members: beatles,
		Albums:  albumsBeatles,
	}

	//fmt.Println(bands)

	err := tpl.Execute(os.Stdout, bands)
	if err != nil {
		log.Fatalln(err)
	}

	//paste it on the .gohtml file

}
