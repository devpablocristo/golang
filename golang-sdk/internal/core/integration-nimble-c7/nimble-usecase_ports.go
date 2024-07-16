package core

import (
	nimble "github.com/devpablocristo/qh/events/internal/core/integration-nimble-c7/nimble"
)

type NimbleUseCasePort interface {
	ProcessOrder(nimble.Order) error
}
