package application

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	uuid "github.com/google/uuid"

	port "github.com/devpablocristo/99minutos/order-manager/internal/application/port"
	domain "github.com/devpablocristo/99minutos/order-manager/internal/domain"
)

type OrderManager struct {
	storage port.OrderRepo
	mux     sync.Mutex
}

func NewOrderManager(st port.OrderRepo) port.OrderManager {
	return &OrderManager{
		storage: st,
		mux:     sync.Mutex{},
	}
}

func (om *OrderManager) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	err := validateCoordinates(order.OrigAddress.Coords)
	if err != nil {
		//log.Println("invalid origin coordinates - " + err.Error())
		return nil, errors.New("invalid origin coordinates - " + err.Error())
	}

	err = validateCoordinates(order.DestAddress.Coords)
	if err != nil {
		//log.Println("invalid destiny coordinates - " + err.Error())
		return nil, errors.New("invalid destiny coordinates - " + err.Error())
	}

	err = validateProducts(order.Products)
	if err != nil {
		//log.Println("invalid products" + err.Error())
		return nil, errors.New("invalid products - " + err.Error())
	}

	totalWeight := calculateTotalWeight(order.Products)
	if totalWeight > 25 {
		//log.Println("order not available with standard service contact the company to make a special arrangement")
		return nil, errors.New("order not available with standard service contact the company to make a special arrangement")
	}

	order.UUID = uuid.New().String()
	order.PackageType = packageType(totalWeight)
	order.TotalWeight = totalWeight
	order.Status = domain.CREATED
	order.CreatedAt = time.Now().Unix()

	om.storage.Create(ctx, order)

	return order, nil
}

func validateCoordinates(coords domain.Coords) error {
	if coords.Latitude < -90 || coords.Latitude > 90 {
		return fmt.Errorf("invalid latitude: %f", coords.Latitude)
	}
	if coords.Longitude < -180 || coords.Longitude > 180 {
		return fmt.Errorf("invalid longitude: %f", coords.Longitude)
	}
	return nil
}

func validateProducts(products []domain.Product) error {
	for _, p := range products {
		if p.Quantity <= 0 {
			return errors.New("invalid quantity")
		}
		if p.Weight <= 0 {
			return errors.New("invalid weight")
		}
		if p.Size <= 0 {
			return errors.New("invalid size")
		}
	}
	return nil
}

func calculateTotalWeight(products []domain.Product) float32 {
	totalWeight := float32(0)
	for _, p := range products {
		totalWeight += p.Weight * float32(p.Quantity)
	}
	return totalWeight
}

func packageType(weight float32) int16 {
	if weight <= 5.0 {
		return domain.SM
	}
	if weight <= 15.0 {
		return domain.MD
	}
	return domain.LG
}
