package main

import (
	"log"

	is "github.com/devpablocristo/qh/events/pkg/init-setup"

	arts "github.com/devpablocristo/qh/events/cmd/rest/auth/routes"
	mon "github.com/devpablocristo/qh/events/cmd/rest/monitoring"
	nrts "github.com/devpablocristo/qh/events/cmd/rest/nimble-cin7/routes"
	urts "github.com/devpablocristo/qh/events/cmd/rest/user/routes"
	csd "github.com/devpablocristo/qh/events/internal/platform/cassandra"
	cnsl "github.com/devpablocristo/qh/events/internal/platform/consul"
	gin "github.com/devpablocristo/qh/events/internal/platform/gin"
	gmw "github.com/devpablocristo/qh/events/internal/platform/go-micro-web"
	mysql "github.com/devpablocristo/qh/events/internal/platform/mysql"
	pgts "github.com/devpablocristo/qh/events/internal/platform/postgresql/pqxpool"
	redis "github.com/devpablocristo/qh/events/internal/platform/redis"
	stg "github.com/devpablocristo/qh/events/internal/platform/stage"
)

func main() {
	if err := is.InitSetup(); err != nil {
		log.Fatalf("Error setting up configurations: %v", err)
	}
	is.LogInfo("Application started with JWT secret key: %s", is.GetJWTSecretKey())
	is.MicroLogInfo("Starting application...")

	// TODO: Probar stage
	// stage
	if _, err := stg.NewStageInstance(); err != nil {
		is.MicroLogError("error initializing Stage: %v", err)
	}

	if _, err := cnsl.NewConsulInstance(); err != nil {
		is.MicroLogError("error initializing Consul: %v", err)
	}

	// TODO: Probar go micro
	ms, err := gmw.NewGoMicroInstance()
	if err != nil {
		is.MicroLogError("error initializing Go Micro: %v", err)
	}

	if _, err = pgts.NewPostgreSQLInstance(); err != nil {
		is.MicroLogError("error initializing PostgresSQL: %v", err)
	}

	if _, err = csd.NewCassandraInstance(); err != nil {
		is.MicroLogError("error initializing Canssandra: %v", err)
	}

	if _, err = redis.NewRedisInstance(); err != nil {
		is.MicroLogError("error initializing Redis: %v", err)
	}

	if _, err = mysql.NewMySQLInstance(); err != nil {
		is.MicroLogError("error initializing MySQL: %v", err)
	}

	ginInst, err := gin.NewGinInstance()
	if err != nil {
		is.MicroLogError("error initializing Gin: %v", err)
	}

	r := ginInst.GetRouter()

	nrts.NimRoutes(r)
	urts.UserRoutes(r)
	arts.AuthRoutes(r)

	mon.MonitoringRestAPI(ginInst, ms)

	if err := ginInst.RunServer(); err != nil {
		is.MicroLogError("error starting Gin server: %v", err)
	}

	// TODO: pyroscope
	// TODO: probar pprof
	// TODO: implmentar multi-tenancy

	// TODO: implementar context, como este ejemplo:
	// Crea un contexto con un tiempo de espera de 50 segundos.
	// El contexto se utiliza para definir cuánto tiempo debe esperar el programa
	// antes de que una operación de MongoDB se considere fallida. defer cancel()
	// garantiza que el contexto se cancele cuando la función main termine.
	// ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	// defer cancel()

	// TODO: implementar rabbitmq users
	// rabbitmqURI := os.Getenv("RABBITMQ_URI")
	// if rabbitmqURI == "" {
	// 	log.Fatal("RABBITMQ_URI no está definido en las variables de entorno")
	// }
	// log.Println("RABBITMQ_URI:", rabbitmqURI)

	// log.Println("Initializing repositories and use cases...")
	// repository := usr.NewRepository()
	// usecase := ucs.NewUseCase(repository)
	// restHandler := hdl.NewRestHandler(usecase)

	// log.Println("Initializing RabbitMQ...")
	// rabbitMqManager, err := hdl.NewRabbitHandler(rabbitmqURI, "myQueue", usecase)
	// if err != nil {
	// 	log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	// }
	// defer func() {
	// 	log.Println("Closing RabbitMQ connection...")
	// 	rabbitMqManager.Close()
	// }()

	// log.Println("Starting to consume messages...")
	// err = rabbitMqManager.ConsumeMessages()
	// if err != nil {
	// 	log.Fatalf("Failed to start consuming messages: %s", err)
	// }

	// r := gin.Default()
	// log.Println("Setting up routes...")
	// r.GET("/users/:id", restHandler.GetUser)
	// r.POST("/users", restHandler.CreateUser)
	// r.PUT("/users/:id", restHandler.UpdateUser)
	// r.DELETE("/users/:id", restHandler.DeleteUser)
	// r.GET("/users", restHandler.ListUsers)
	// r.GET("/", restHandler.HelloWorld)

	// go func() {
	// 	log.Println("Running server on port 8080...")
	// 	if err := r.Run(":8080"); err != nil {
	// 		log.Fatalf("Failed to run server: %s", err)
	// 	}
	// }()

	// // Manejar señales del sistema para una finalización ordenada
	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// <-quit
	// log.Println("Shutting down server...")

	// TODO: rating
	// usecase := ucs.NewUseCase(dep.Repository, dep.ApiClient)
	// handler := hdl.NewRestHandler(usecase)

	// r := gin.Default()

	// v1 := r.Group("/api/v1")
	// {
	// 	v1.GET("/ltp", handler.GetLTP)
	// }

	// log.Println("Running server on port " + dep.RouterPort + "...")
	// if err := r.Run(":" + dep.RouterPort); err != nil {
	// 	log.Fatalf("Failed to run server: %s", err)
	// }

	// return r

}
