package core

import (
	c7 "github.com/devpablocristo/qh/events/internal/core/integration-nimble-c7/c7"
)

type Cin7UseCase struct {
	repo c7.RedisPort
}

func NewCin7UseCase(repo c7.RedisPort) Cin7UseCasePort {
	return &Cin7UseCase{
		repo: repo,
	}
}

func (uc *Cin7UseCase) UpdateShipment(shipment c7.Shipment) error {
	return uc.repo.SaveShipment(shipment)
}
