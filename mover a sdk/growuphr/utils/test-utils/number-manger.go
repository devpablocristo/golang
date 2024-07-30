package testutils

import (
	"fmt"
	"net/http"
	"time"

	cmsapi "github.com/devpablocristo/growuphr/internal/commons/api"
	domain "github.com/devpablocristo/growuphr/number-manager/domain"
)

var (
	ResNum0 = &domain.ReservedNumber{
		UUID:      "0",
		Number:    &num0,
		User:      &usr0,
		CreatedAt: time.Time{},
	}

	ResNum1 = &domain.ReservedNumber{
		UUID:      "1",
		Number:    &num1,
		User:      &usr1,
		CreatedAt: time.Time{},
	}

	ResNum2 = &domain.ReservedNumber{
		UUID:      "2",
		Number:    &num2,
		User:      &usr2,
		CreatedAt: time.Time{},
	}

	num0 = domain.Number{
		UUID:      "0",
		Number:    0,
		CreatedAt: time.Time{},
	}

	usr0 = domain.User{
		UUID:      "0",
		Username:  "client0",
		CreatedAt: time.Time{},
	}

	num1 = domain.Number{
		UUID:      "1",
		Number:    1,
		CreatedAt: time.Time{},
	}

	usr1 = domain.User{
		UUID:      "1",
		Username:  "client1",
		CreatedAt: time.Time{},
	}

	num2 = domain.Number{
		UUID:      "2",
		Number:    2,
		CreatedAt: time.Time{},
	}

	usr2 = domain.User{
		UUID:      "2",
		Username:  "client2",
		CreatedAt: time.Time{},
	}

	ErrorZeroNumber   = cmsapi.NewAPIError(http.StatusOK, "taken-number", "AddReserveNumber", "application", fmt.Errorf("number '%v' is invalid, try with another one", ResNum0.Number.Number))
	ErrorTakenNumber  = cmsapi.NewAPIError(http.StatusOK, "taken-number", "AddReserveNumber", "application", fmt.Errorf("number '%v' already reserved, try with another one", ResNum1.Number.Number))
	ErrorExistingUser = cmsapi.NewAPIError(http.StatusOK, "taken-number", "AddReserveNumber", "application", fmt.Errorf("user '%v' has already reserved the number '%v'", ResNum1.User.Username, ResNum1.Number.Number))
)
