package application

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	uuid "github.com/google/uuid"

	cmsapi "github.com/devpablocristo/growuphr/internal/commons/api"
	port "github.com/devpablocristo/growuphr/number-manager/application/port"
	domain "github.com/devpablocristo/growuphr/number-manager/domain"
)

type NumberManager struct {
	storage port.Storage
	mux     sync.Mutex
}

func NewNumberManager(st port.Storage) *NumberManager {
	return &NumberManager{
		storage: st,
		mux:     sync.Mutex{},
	}
}

func (nm *NumberManager) AddReserveNumber(ctx context.Context, newResNum *domain.ReservedNumber) error {
	nm.mux.Lock()
	defer nm.mux.Unlock()

	if newResNum.Number.Number == 0 {
		eMsg := fmt.Errorf("number '%v' is invalid, try with another one", newResNum.Number.Number)
		err := cmsapi.NewAPIError(200, "taken-number", "AddReserveNumber", "application", eMsg)
		log.Println(err.Error())
		return err
	}

	rn, foundUsr := nm.storage.CheckForUsername(ctx, newResNum.User.Username)
	if foundUsr {
		eMsg := fmt.Errorf("user '%v' has already reserved the number '%v'", rn.User.Username, rn.Number.Number)
		err := cmsapi.NewAPIError(http.StatusOK, "taken-number", "AddReserveNumber", "application", eMsg)
		log.Println(err.Error())
		return err
	}

	rn, foundNum := nm.storage.CheckForNumber(ctx, newResNum.Number.Number)
	if foundNum {
		eMsg := fmt.Errorf("number '%v' already reserved, try with another one", rn.Number.Number)
		err := cmsapi.NewAPIError(http.StatusOK, "taken-number", "AddReserveNumber", "application", eMsg)
		log.Println(err.Error())
		return err
	}

	newResNum.UUID = uuid.New().String()
	newResNum.CreatedAt = time.Now()

	newResNum.Number.UUID = uuid.New().String()
	newResNum.Number.CreatedAt = newResNum.CreatedAt

	newResNum.User.UUID = uuid.New().String()
	newResNum.User.CreatedAt = newResNum.CreatedAt

	err := nm.storage.Create(ctx, newResNum)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (nm *NumberManager) ReservedNumbers(ctx context.Context) (map[string]*domain.ReservedNumber, error) {
	nm.mux.Lock()
	defer nm.mux.Unlock()

	nm.storage.List(ctx)
	list := nm.storage.List(ctx)
	if len(list) == 0 {
		eMsg := errors.New("empty list")
		err := cmsapi.NewAPIError(http.StatusOK, "taken-number", "ReservedNumbers", "application", eMsg)
		log.Println(err.Error())
		return nil, err
	}
	return list, nil
}
