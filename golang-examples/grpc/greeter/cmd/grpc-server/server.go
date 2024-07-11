package main

import (
	"log"

	grpchandler "greet/internal/adapters/driver/handlers/grpc"
	service "greet/internal/service"
)

// Main function to start the server
func main() {

	// a child context is being created to limit the duration a operation,
	// since when calling the cancellation function,
	// any routine waiting for the closed context will immediately return with an error.
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// repo()
	// app(repo)
	// handler(app)

	//greetService := service.NewGreetServer()

	//db, err := eventdb.MysqlConn()
	// db, err := eventdb.GormConn()
	// if err != nil {
	// 	panic(err)
	// }

	// repo := greetdb.NewEventRepository(db)
	// greetService := service.NewGreetService(repo)
	greetService := service.NewGreetService()
	grpcHandler := grpchandler.NewGreetHandler(greetService)

	err := grpcConn(grpcHandler)
	if err != nil {
		log.Fatalln(err)
	}
}
