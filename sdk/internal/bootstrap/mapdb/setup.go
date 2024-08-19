package mapdbsetup

import (
	"errors"

	mapdb "github.com/devpablocristo/golang/sdk/pkg/mapdb/std"
	portspkg "github.com/devpablocristo/golang/sdk/pkg/mapdb/std/portspkg"
)

func NewMapDbInstance() (portspkg.MapDbClient, error) {
	// Crear una nueva configuración para MapDb
	var config portspkg.MapDbConfig

	// Inicializar el cliente de MapDb con la configuración
	if err := mapdb.InitializeMapDbClient(config); err != nil {
		return nil, err
	}

	// Obtener la instancia de MapDb
	instance, err := mapdb.GetMapDbInstance()
	if err != nil {
		return nil, errors.New("failed to get MapDb instance")
	}

	return instance, nil
}

// Generics
// // NewMapDbInstance inicializa y devuelve una instancia de MapDb con la configuración proporcionada.
// func NewMapDbInstance[T any](prepopulate bool, data []T) (portspkg.MapDbClient[T], error) {
// 	// Crear una instancia de configuración de MapDb usando NewMapDbConfig
// 	config := mapdb.NewMapDbConfig(prepopulate, data)

// 	// Validar la configuración antes de inicializar el cliente
// 	if err := config.Validate(); err != nil {
// 		return nil, err
// 	}

// 	// Inicializar el cliente de MapDb con la configuración
// 	if err := mapdb.InitializeMapDbClient(config); err != nil {
// 		return nil, err
// 	}

// 	// Obtener la instancia de MapDb
// 	return mapdb.GetMapDbInstance[T]()
// }
