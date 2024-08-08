package core

import (
	nimble "github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/nimble"
)

type NimbleUseCasesPort interface {
	ProcessOrder(nimble.Order) error
}

type nimbleUseCases struct {
	repo         nimble.CachePort
	cin7UseCases Cin7UseCasesPort
}

func NewNimbleUseCases(repo nimble.CachePort, cin7UseCases Cin7UseCasesPort) NimbleUseCasesPort {
	return &nimbleUseCases{
		repo:         repo,
		cin7UseCases: cin7UseCases,
	}
}

func (uc *nimbleUseCases) ProcessOrder(order nimble.Order) error {
	// Transforma la orden de Nimble a un formato de envío de Cin7
	shipment, err := uc.repo.CreateShipment(order)
	if err != nil {
		return err
	}
	// Llama al caso de uso de Cin7 para actualizar el envío
	return uc.cin7UseCases.UpdateShipment(shipment)
}
