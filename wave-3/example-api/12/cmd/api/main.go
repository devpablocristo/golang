package main

import (
	"log"

	cmsenv "items/common/env"
	controller "items/internal/adapters/controller"
	repository "items/internal/adapters/repository"
	mysql "items/internal/infra/mysql"
	web "items/internal/infra/web"
	usecase "items/internal/usecase"
)

func main() {

	cmsenv.LoadEnv()

	//MySQL
	conn, err := mysql.GetConnectionDB()
	if err != nil {
		log.Fatalln(err)
	}

	m := repository.NewMySQLItemRepository(conn)
	r := repository.NewRepository()
	u := usecase.NewItemUsecase(r, m)
	h := controller.NewController(u)

	// se mueven ls rutas a otro archivo
	err = web.NewHTTPServer(h)
	if err != nil {
		log.Fatalln(err)
	}
}
