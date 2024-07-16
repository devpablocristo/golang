package core

import (
	nimble "github.com/devpablocristo/qh/events/internal/core/integration-nimble-cin7/nimble"
)

type NimbleUseCasePort interface {
	ProcessOrder(nimble.Order) error
}
