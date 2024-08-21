package mapdbsetup

import (
	"errors"

	mapdb "github.com/devpablocristo/golang/sdk/pkg/mapdb/std"
	portspkg "github.com/devpablocristo/golang/sdk/pkg/mapdb/std/portspkg"
)

func NewMapDbInstance() (portspkg.MapDbClient, error) {
	if err := mapdb.InitializeMapDbClient(); err != nil {
		return nil, err
	}

	instance, err := mapdb.GetMapDbInstance()
	if err != nil {
		return nil, errors.New("failed to get MapDb instance")
	}

	return instance, nil
}
