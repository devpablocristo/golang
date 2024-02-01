package main

import (
	"log"

	//"items/internal/adapters/repository/inmemoryr"
	//"items/internal/adapters/repository/mysqlr"

	"items/internal/adapters/handler"
	"items/internal/adapters/repository/mysqlr"
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
	r := mysqlr.NewMySQLRepository(conn)
	//

	//r := inmemoryr.NewInMemory()
	u := usecase.NewItemUsecase(r)
	h := handler.NewHandler(u)

	web.NewHTTPServer(h)
}
