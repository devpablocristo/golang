package core

import (
	cin7 "github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/cin7"
)

type Cin7UseCases interface {
	UpdateShipment(cin7.Shipment) error
}

type cin7UseCases struct {
	repo cin7.CacheRepository
}

func NewCin7UseCases(repo cin7.CacheRepository) Cin7UseCases {
	return &cin7UseCases{
		repo: repo,
	}
}

func (uc *cin7UseCases) UpdateShipment(shipment cin7.Shipment) error {
	return uc.repo.SaveShipment(shipment)
}
