package main

import (
	"log"

	"items/internal/adapters/handler"
	mysqlr "items/internal/adapters/repository/mysql"

	//"items/internal/adapters/repository/inmemory"
	"items/internal/platform/mysql"
	"items/internal/platform/web"
	"items/internal/usecase"
)

func main() {
	//MySQL
	conn, err := mysql.GetConnectionDB()
	if err != nil {
		log.Fatalln(err)
	}

	m := mysqlr.NewMySQLRepository(conn)

	//r := repository.NewRepository()
	u := usecase.NewItemUsecase(m)
	h := handler.NewHandler(u)

	err = web.NewHTTPServer(h)
	if err != nil {
		log.Fatalln(err)
	}
}
