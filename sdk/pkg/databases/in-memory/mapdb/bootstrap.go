package pkgmapdb

import (
	ports "github.com/devpablocristo/golang/sdk/pkg/databases/in-memory/mapdb/ports"
)

func Boostrap() (ports.Service, error) {
	NewService()
	return GetInstance()
}
