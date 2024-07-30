package api

import (
	"sync"

	application "github.com/devpablocristo/blankfactor/event-service/internal/application"
	//eventdb "github.com/devpablocristo/blankfactor/event-service/internal/infrastructure/driven-adapter/repository/mysql"
	eventdb "github.com/devpablocristo/blankfactor/event-service/internal/infrastructure/driven-adapter/repository/gorm"
	handler "github.com/devpablocristo/blankfactor/event-service/internal/infrastructure/driver-adapter/handler"
)

func StartApi(wg *sync.WaitGroup, port string) {
	//db, err := eventdb.MysqlConn()
	db, err := eventdb.GormConn()
	if err != nil {
		panic(err)
	}

	rpo := eventdb.NewEventRepository(db)
	esv := application.NewEventService(rpo)
	han := handler.NewHandler(esv)
	rou := Router(han)

	HttpServer(port, rou)
}
