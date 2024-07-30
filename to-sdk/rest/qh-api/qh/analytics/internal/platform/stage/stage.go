package stage

type Stage int

const (
	Development Stage = iota
	Beta
	Production
)

var stageNames = map[Stage]string{
	Development: "DEV",
	Beta:        "BETA",
	Production:  "PROD",
}

func (s Stage) String() string {
	return stageNames[s]
}

func GetFromString(str string) Stage {
	for key, value := range stageNames {
		if value == str {
			return key
		}
	}
	return Development
}
