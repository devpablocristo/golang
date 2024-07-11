package main

import (
	"log"

	ctrl "github.com/mercadolibre/items/internal/adapters/controller"
	"github.com/mercadolibre/items/internal/adapters/repository"
	"github.com/mercadolibre/items/internal/infra/mysql"
	"github.com/mercadolibre/items/internal/infra/web"
	"github.com/mercadolibre/items/internal/usecase"
)

func main() {
	//MySQL
	conn, err := mysql.GetConnectionDB()
	if err != nil {
		log.Fatalln(err)
	}

  itemRepository := repository.NewMySQLItemRepository(conn)
	//itemRepository := repository.NewItemRepository()
	itemUsecase := usecase.NewItemUsecase(itemRepository)
	itemController := ctrl.NewItemController(itemUsecase)

	if err := web.NewHTTPServer(itemController); err != nil {

		log.Fatalln(err)
	}
}
