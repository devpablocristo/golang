package main

import (
	"log"

	ctrl "github.com/devpablocristo/golang/06-projects/items/gin/items-final/internal/adapters/handler"
	repository "github.com/devpablocristo/golang/06-projects/items/gin/items-final/internal/adapters/repository"
	mysql "github.com/devpablocristo/golang/06-projects/items/gin/items-final/internal/infra/mysql"
	web "github.com/devpablocristo/golang/06-projects/items/gin/items-final/internal/infra/web"
	usecase "github.com/devpablocristo/golang/06-projects/items/gin/items-final/internal/usecase"
)

func main() {
	//MySQL
	conn, err := mysql.GetConnectionDB()
	if err != nil {
		log.Fatalln(err)
	}

	itemRepository := repository.NewMySQLRepository(conn)
	//itemRepository := repository.NewItemRepository()
	itemUsecase := usecase.NewItemUsecase(itemRepository)
	itemHandler := ctrl.NewItemHandler(itemUsecase)

	if err := web.NewHTTPServer(itemHandler); err != nil {

		log.Fatalln(err)
	}
}
