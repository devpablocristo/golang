package main

import (
	"log"

	handler "items/internal/adapters/handler"
	repository "items/internal/adapters/repository"
	"items/internal/infra/mysql"
	web "items/internal/infra/web"
	usecase "items/internal/usecase"
)

func main() {

	//MySQL
	conn, err := mysql.GetConnectionDB()
	if err != nil {
		log.Fatalln(err)
	}

	m := repository.NewMySQLItemRepository(conn)
	r := repository.NewRepository()
	u := usecase.NewItemUsecase(r, m)
	h := handler.NewHandler(u)

	// se mueven ls rutas a otro archivo
	err = web.NewHTTPServer(h)
	if err != nil {
		log.Fatalln(err)
	}
}
