package core

import (
	cin7 "github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/cin7"
)

type Cin7UseCasePort interface {
	UpdateShipment(cin7.Shipment) error
}

type cin7UseCase struct {
	repo cin7.CachePort
}

func NewCin7UseCase(repo cin7.CachePort) Cin7UseCasePort {
	return &cin7UseCase{
		repo: repo,
	}
}

func (uc *cin7UseCase) UpdateShipment(shipment cin7.Shipment) error {
	return uc.repo.SaveShipment(shipment)
}
