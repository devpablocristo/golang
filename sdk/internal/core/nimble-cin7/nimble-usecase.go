package core

import (
	nimble "github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7/nimble"
)

type NimbleUseCase struct {
	repo        nimble.RedisPort
	cin7UseCase Cin7UseCasePort
}

func NewNimbleUseCase(repo nimble.RedisPort, cin7UseCase Cin7UseCasePort) NimbleUseCasePort {
	return &NimbleUseCase{
		repo:        repo,
		cin7UseCase: cin7UseCase,
	}
}

func (uc *NimbleUseCase) ProcessOrder(order nimble.Order) error {
	// Transforma la orden de Nimble a un formato de envío de Cin7
	shipment, err := uc.repo.CreateShipment(order)
	if err != nil {
		return err
	}
	// Llama al caso de uso de Cin7 para actualizar el envío
	return uc.cin7UseCase.UpdateShipment(shipment)
}
