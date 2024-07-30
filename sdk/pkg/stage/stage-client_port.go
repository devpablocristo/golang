package stage

type StageClientPort interface {
	String(s Stage) string
	GetFromString(str string) Stage
	GetFromCamelCase(str string) Stage
	GetFromKebabCase(str string) Stage
}
