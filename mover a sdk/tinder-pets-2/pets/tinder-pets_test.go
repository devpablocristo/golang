package tinderpets_test

import (
	"log"
	"testing"

	tinder "github.com/devpablocristo/go-concepts/std-lib/testing/tinder-pets"
)

// func TestAll(t *testing.T) {

// 	t.Run("fail", TestCreateProfile)

// }

// func TestCreateProfile(t *testing.T) {

// 	ds := make([]tinder.Dog, 1)
// 	dogTest := tinder.Dog{
// 		Race:  "",
// 		Size:  "",
// 		Name:  "",
// 		Sex:   "",
// 		Color: "",
// 	}
// 	ds[0] = dogTest

// 	want := tinder.Profile{
// 		Email:    "",
// 		PassWord: "",
// 		Dog:      ds,
// 	}

// 	get := tinder.CreateProfile(want)

// 	if !reflect.DeepEqual(get[0], want) {
// 		t.Errorf("I want %q", want)
// 		t.Errorf("I get %q", get[0])
// 	} else {
// 		log.Println("All ok!")
// 		log.Println("want: ", want)
// 		log.Println("get: ", get[0])
// 	}
// }

func TestMatchPets(t *testing.T) {

	pet1 := tinder.Pet{
		ID:     1,
		Specie: "cat",
		Race:   "tiger",
		Size:   "medium",
		Name:   "Tori",
		Sex:    "male",
		Color:  "blonde",
	}
	pet1.MatchsIDs = append(pet1.MatchsIDs, 12)

	pet2 := tinder.Pet{
		ID:     12,
		Specie: "cat",
		Race:   "tiger",
		Size:   "small",
		Name:   "Juliet",
		Sex:    "female",
		Color:  "white",
	}
	pet2.MatchsIDs = append(pet2.MatchsIDs, 1)

	if pet1.MatchsIDs[0] != pet2.ID {
		t.Errorf("Want %q: ", pet1)
		t.Errorf("Match with: %q", pet2)

	} else if pet2.MatchsIDs[0] != pet1.ID {
		t.Errorf("Want %q match with %q", pet1, pet2)
	} else {
		log.Printf("Success!!\n")
		log.Printf("%v match with \n", pet1)
		log.Printf("%v\n", pet2)
	}

}
