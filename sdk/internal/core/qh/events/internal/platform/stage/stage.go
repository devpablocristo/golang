package stage

type Stage int

const (
	Unknown Stage = iota
	Development
	Beta
	Production
)

var stageNames = map[Stage]string{
	Unknown:     "UNKNOWN",
	Development: "DEV",
	Beta:        "BETA",
	Production:  "PROD",
}

func (s Stage) String() string {
	if name, exists := stageNames[s]; exists {
		return name
	}
	return "UNKNOWN"
}

func GetFromString(str string) Stage {
	for key, value := range stageNames {
		if value == str {
			return key
		}
	}
	return Unknown
}
