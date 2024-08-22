package pkgmapdbports

type Service interface {
	GetDb() map[string]any
}
