package main

import (
	"log"

	is "github.com/devpablocristo/golang/sdk/pkg/init-setup"

	//arts "github.com/devpablocristo/golang/sdk/cmd/rest/auth/routes"
	//mon "github.com/devpablocristo/golang/sdk/cmd/rest/monitoring"
	//nrts "github.com/devpablocristo/golang/sdk/cmd/rest/nimble-cin7/routes"

	//csd "github.com/devpablocristo/golang/sdk/internal/platform/cassandra"
	//cnsl "github.com/devpablocristo/golang/sdk/internal/platform/consul"
	gin "github.com/devpablocristo/golang/sdk/internal/platform/gin"
	//gmw "github.com/devpablocristo/golang/sdk/internal/platform/go-micro-web"

	//mysql "github.com/devpablocristo/golang/sdk/internal/platform/mysql"
	//pgts "github.com/devpablocristo/golang/sdk/internal/platform/postgresql/pqxpool"
	//redis "github.com/devpablocristo/golang/sdk/internal/platform/redis"
	//stg "github.com/devpablocristo/golang/sdk/internal/platform/stage"

	user "github.com/devpablocristo/golang/sdk/cmd/rest/user/routes"
)

func main() {
	if err := is.InitSetup(); err != nil {
		log.Fatalf("Error setting up configurations: %v", err)
	}
	is.LogInfo("Application started with JWT secret key: %s", is.GetJWTSecretKey())
	is.MicroLogInfo("Starting application...")

	// TODO: Probar stage
	// stage
	// if _, err := stg.NewStageInstance(); err != nil {
	// 	is.MicroLogError("error initializing Stage: %v", err)
	// }

	// if _, err := cnsl.NewConsulInstance(); err != nil {
	// 	is.MicroLogError("error initializing Consul: %v", err)
	// }

	// TODO: Probar go micro
	// ms, err := gmw.NewGoMicroInstance()
	// if err != nil {
	// 	is.MicroLogError("error initializing Go Micro: %v", err)
	// }

	// if _, err = pgts.NewPostgreSQLInstance(); err != nil {
	// 	is.MicroLogError("error initializing PostgresSQL: %v", err)
	// }

	// if _, err = csd.NewCassandraInstance(); err != nil {
	// 	is.MicroLogError("error initializing Canssandra: %v", err)
	// }

	// if _, err = redis.NewRedisInstance(); err != nil {
	// 	is.MicroLogError("error initializing Redis: %v", err)
	// }

	// if _, err = mysql.NewMySQLInstance(); err != nil {
	// 	is.MicroLogError("error initializing MySQL: %v", err)
	// }

	ginInst, err := gin.NewGinInstance()
	if err != nil {
		is.MicroLogError("error initializing Gin: %v", err)
	}

	r := ginInst.GetRouter()

	//nrts.NimRoutes(r)
	user.Routes(r)
	//arts.AuthRoutes(r)

	//mon.MonitoringRestAPI(ginInst, ms)

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

	// TODO: ambassadors
	// database.Connect()
	// database.AutoMigrate()
	// database.SetupRedis()
	// database.SetupCacheChannel()

	// app := fiber.New()

	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true,
	// }))

	// routes.Setup(app)

	// app.Listen(":8000")

	// TODO: ambassadors orders
	// "ambassador/src/database"
	// "ambassador/src/models"
	// "github.com/bxcodec/faker/v3"
	// "math/rand"

	// database.Connect()

	// for i := 0; i < 30; i++ {
	// 	var orderItems []models.OrderItem

	// 	for j := 0; j < rand.Intn(5); j++ {
	// 		price := float64(rand.Intn(90) + 10)
	// 		qty := uint(rand.Intn(5))

	// 		orderItems = append(orderItems, models.OrderItem{
	// 			ProductTitle:      faker.Word(),
	// 			Price:             price,
	// 			Quantity:          qty,
	// 			AdminRevenue:      0.9 * price * float64(qty),
	// 			AmbassadorRevenue: 0.1 * price * float64(qty),
	// 		})
	// 	}

	// 	database.DB.Create(&models.Order{
	// 		UserId:          uint(rand.Intn(30) + 1),
	// 		Code:            faker.Username(),
	// 		AmbassadorEmail: faker.Email(),
	// 		FirstName:       faker.FirstName(),
	// 		LastName:        faker.LastName(),
	// 		Email:           faker.Email(),
	// 		Complete:        true,
	// 		OrderItems:      orderItems,
	// 	})
	// }

	//TODO: ambassadors productcs
	// 	"ambassador/src/database"
	// "ambassador/src/models"
	// "github.com/bxcodec/faker/v3"
	// "math/rand"

	// database.Connect()

	// for i := 0; i < 30; i++ {
	// 	product := models.Product{
	// 		Title:       faker.Username(),
	// 		Description: faker.Username(),
	// 		Image:       faker.URL(),
	// 		Price:       float64(rand.Intn(90) + 10),
	// 	}

	// 	database.DB.Create(&product)
	// }

	// TODO: ambassdors users

	// "ambassador/src/database"
	// "ambassador/src/models"
	// "github.com/bxcodec/faker/v3"
	// 	database.Connect()

	// for i := 0; i < 30; i++ {
	// 	ambassador := models.User{
	// 		FirstName:    faker.FirstName(),
	// 		LastName:     faker.LastName(),
	// 		Email:        faker.Email(),
	// 		IsAmbassador: true,
	// 	}

	// 	ambassador.SetPassword("1234")

	// 	database.DB.Create(&ambassador)
	// }

	// TODO: ambassadors ranking
	// 	"ambassador/src/database"
	// "ambassador/src/models"
	// "context"
	// "github.com/go-redis/redis/v8"

	// database.Connect()
	// database.SetupRedis()

	// ctx := context.Background()

	// var users []models.User

	// database.DB.Find(&users, models.User{
	// 	IsAmbassador: true,
	// })

	// for _, user := range users {
	// 	ambassador := models.Ambassador(user)
	// 	ambassador.CalculateRevenue(database.DB)

	// 	database.Cache.ZAdd(ctx, "rankings", &redis.Z{
	// 		Score:  *ambassador.Revenue,
	// 		Member: user.Name(),
	// 	})
	// }

	// TODO articlehub
	// func main() {
	// 	// Set Gin to production mode
	// 	gin.SetMode(gin.ReleaseMode)

	// 	// Set the router as the default one provided by Gin
	// 	router = gin.Default()

	// 	// Process the templates at the start so that they don't have to be loaded
	// 	// from the disk again. This makes serving HTML pages very fast.
	// 	router.LoadHTMLGlob("templates/*")

	// 	// Initialize the routes
	// 	initializeRoutes()

	// 	// Start serving the application
	// 	router.Run()
	// }

	// // Render one of HTML, JSON or CSV based on the 'Accept' header of the request
	// // If the header doesn't specify this, HTML is rendered, provided that
	// // the template name is present
	// func render(c *gin.Context, data gin.H, templateName string) {
	// 	loggedInInterface, _ := c.Get("is_logged_in")
	// 	data["is_logged_in"] = loggedInInterface.(bool)

	// 	switch c.Request.Header.Get("Accept") {
	// 	case "application/json":
	// 		// Respond with JSON
	// 		c.JSON(http.StatusOK, data["payload"])
	// 	case "application/xml":
	// 		// Respond with XML
	// 		c.XML(http.StatusOK, data["payload"])
	// 	default:
	// 		// Respond with HTML
	// 		c.HTML(http.StatusOK, templateName, data)
	// 	}
	// }

}
