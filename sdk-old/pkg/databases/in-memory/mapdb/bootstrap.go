package sdkmapdb

import (
	defs "github.com/devpablocristo/golang/sdk/pkg/databases/in-memory/mapdb/defs"
)

func Boostrap() defs.Repository {
	return newRepository()
}
