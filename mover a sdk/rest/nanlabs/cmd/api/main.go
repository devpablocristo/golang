package main

import (
	"os"

	application "github.com/devpablocristo/nanlabs/application"
	//domain "github.com/devpablocristo/nanlabs/domain"
	mapdb "github.com/devpablocristo/nanlabs/infrastructure/driven-adapter/repository/mapdb"
	trelloservice "github.com/devpablocristo/nanlabs/infrastructure/driven-adapter/trello"
	chihandler "github.com/devpablocristo/nanlabs/infrastructure/driver-adapter/handler/chi"
)

const defaultPort = "8080"

var (
	TRELLO_KEY   string = "64f8e58c56392537750ddb333e2ed257"
	TRELLO_TOKEN string = "ATTA0f11e8604137ea1b2222b196c12e50a175ced7a86cabdd0a0c9722f9b22bb916DF28BB37"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mdb := mapdb.NewMapDB()
	tse := trelloservice.NewTrelloService(TRELLO_KEY, TRELLO_TOKEN)
	tke := application.NewTaskService(mdb, tse)

	han := chihandler.NewChiHandler(tke)
	rou := SetupRouter(han)
	HttpServer(port, rou)
}
