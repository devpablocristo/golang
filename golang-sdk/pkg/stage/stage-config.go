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
