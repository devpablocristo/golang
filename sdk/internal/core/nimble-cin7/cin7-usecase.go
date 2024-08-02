package core

import (
	cin7 "github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/cin7"
)

type Cin7UseCase struct {
	repo cin7.RedisPort
}

func NewCin7UseCase(repo cin7.RedisPort) Cin7UseCasePort {
	return &Cin7UseCase{
		repo: repo,
	}
}

func (uc *Cin7UseCase) UpdateShipment(shipment cin7.Shipment) error {
	return uc.repo.SaveShipment(shipment)
}
