package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	trello "github.com/adlio/trello"
)

func main() {

	values := map[string]string{
		"idlist": "63c490512d8e4301239ceaa0",
		"key":    "64f8e58c56392537750ddb333e2ed257",
		"token":  "ATTA0f11e8604137ea1b2222b196c12e50a175ced7a86cabdd0a0c9722f9b22bb916DF28BB37",
	}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://api.trello.com/1/cards", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"])

	client := trello.NewClient("64f8e58c56392537750ddb333e2ed257", "ATTA0f11e8604137ea1b2222b196c12e50a175ced7a86cabdd0a0c9722f9b22bb916DF28BB37")

	list, err := client.GetList("63c490512d8e4301239ceaa0", trello.Defaults())
	list.AddCard(&trello.Card{Name: "Card sdsName", Desc: "Card description"}, trello.Defaults())
}
