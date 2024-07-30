package api

import (
	"sync"

	application "github.com/devpablocristo/growuphr/number-manager/application"
	mapdb "github.com/devpablocristo/growuphr/number-manager/infrastructure/driven-adapter/repository/mapdb"
	handler "github.com/devpablocristo/growuphr/number-manager/infrastructure/driver-adapter/handler"
)

func StartApi(wg *sync.WaitGroup, port string) {
	mdb := mapdb.NewMapDB()
	nmr := application.NewNumberManager(mdb)
	han := handler.NewHandler(nmr)
	rou := Router(han)

	HttpServer(port, rou)
}
