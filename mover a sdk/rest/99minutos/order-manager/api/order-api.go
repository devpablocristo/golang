package api

import (
	"context"
	"sync"
	"time"

	application "github.com/devpablocristo/99minutos/order-manager/internal/application"
	domain "github.com/devpablocristo/99minutos/order-manager/internal/domain"
	orderdb "github.com/devpablocristo/99minutos/order-manager/internal/infrastructure/driven-adapter/repository/mapdb/orderdb"
	userdb "github.com/devpablocristo/99minutos/order-manager/internal/infrastructure/driven-adapter/repository/mapdb/userdb"
	valusr "github.com/devpablocristo/99minutos/order-manager/internal/infrastructure/driven-adapter/validate-user"
	handler "github.com/devpablocristo/99minutos/order-manager/internal/infrastructure/driver-adapter/handler"
)

func StartApi(wg *sync.WaitGroup, port string) {

	internal := domain.User{
		UUID:      "1",
		Username:  "internal10",
		Email:     "internal10@99minutos.com",
		Password:  "superPass",
		Role:      domain.INTERNAL,
		CreatedAt: time.Now().Unix(),
	}

	customer := domain.User{
		UUID:      "2",
		Username:  "customer",
		Email:     "customer@mail.com",
		Password:  "12345",
		Role:      domain.CUSTOMER,
		CreatedAt: time.Now().Unix(),
	}

	odb := orderdb.NewMapDB()
	udb := userdb.NewMapDB()

	udb.Create(context.Background(), &internal)
	udb.Create(context.Background(), &customer)

	vlu := valusr.NewValidateUser(udb)
	nmr := application.NewOrderManager(odb)
	han := handler.NewHandler(nmr, vlu, udb, odb)
	rou := Router(han)

	HttpServer(port, rou)
}
