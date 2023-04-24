package usecase

import "errors"

// este error es utilizado por mas de paquete por eso esta en este paquete pero es esto
// en realidad no es una buena implementacion por lo que debera mover
var errNotFound = errors.New("not found")
