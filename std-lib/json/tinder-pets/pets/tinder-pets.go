package tinderpets

type Profile struct {
	Email    string
	PassWord string
	Pets     []Pet
}

type Pet struct {
	ID        int
	Specie    string
	Race      string
	Size      string
	Name      string
	Sex       string
	Color     string
	MatchsIDs []int
}

var (
	Profiles []Profile
)

func CreateProfile(p Profile) []Profile {

	var ps []Profile
	ps = append(ps, p)

	return ps
}

func MatchPets(p1 Pet, p2 Pet) {

}
