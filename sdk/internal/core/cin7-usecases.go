package core

import (
	cin7 "github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/cin7"
)

type Cin7UseCasesPort interface {
	UpdateShipment(cin7.Shipment) error
}

type cin7UseCases struct {
	repo cin7.CachePort
}

func NewCin7UseCases(repo cin7.CachePort) Cin7UseCasesPort {
	return &cin7UseCases{
		repo: repo,
	}
}

func (uc *cin7UseCases) UpdateShipment(shipment cin7.Shipment) error {
	return uc.repo.SaveShipment(shipment)
}
