package portspkg

type MapDbConfig[T any] interface {
	GetPrepopulate() bool
	SetPrepopulate(prepopulate bool)
	GetPrepopulateData() []T
	SetPrepopulateData(data []T)
	Validate() error
}

type MapDbClient[T any] interface {
	Initialize(MapDbConfig[T]) error
}
