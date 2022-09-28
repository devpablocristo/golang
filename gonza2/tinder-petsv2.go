package main

import (
	pets "/Users/gebbes/github.com/gebbes/go-concepts/std-lib/json/tinder-pets"
	"strconv"
)

func GetMatches(Profiles []Profile) map[string][]int {
	matches := make(map[string][]int)

	for _, profile := range Profiles {
		for _, pet := range profile.Pets {
			idString := strconv.Itoa(pet.ID)
			matches[idString] = pet.MatchsIDs
		}
	}
	return matches
}

func main() {
	matches = GetMatches(pets.Profiles)
}
