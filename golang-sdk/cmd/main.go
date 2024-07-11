package main

import (
	"log"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	is "github.com/devpablocristo/qh/events/pkg/init-setup"
	mdhw "github.com/devpablocristo/qh/events/pkg/middleware"

	wire "github.com/devpablocristo/qh/events/cmd/rest"
	csd "github.com/devpablocristo/qh/events/internal/platform/cassandra"
	cnsl "github.com/devpablocristo/qh/events/internal/platform/consul"
	gin "github.com/devpablocristo/qh/events/internal/platform/gin"
	gmw "github.com/devpablocristo/qh/events/internal/platform/go-micro-web"
	mysql "github.com/devpablocristo/qh/events/internal/platform/mysql"
	pgts "github.com/devpablocristo/qh/events/internal/platform/postgres"
	redis "github.com/devpablocristo/qh/events/internal/platform/redis"
	stg "github.com/devpablocristo/qh/events/internal/platform/stage"
)



func main() {
	if err := is.InitSetup(); err != nil {
		log.Fatalf("Error setting up configurations: %v", err)
	}
	is.LogInfo("Application started with JWT secret key: %s", is.GetJWTSecretKey())
	is.MicroLogInfo("Starting application...")

	userHandler, err := wire.InitializeUserHandler()
	if err != nil {
		is.MicroLogError("userHandler error: %v", err)
	}
	authHandler, err := wire.InitializeAuthHandler()
	if err != nil {
		is.MicroLogError("authHandler error: %v", err)
	}

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

	api := r.Group("/api/v1")
	{
		api.POST("/login", authHandler.Login)
	}

	secret := "secret"
	user := r.Group("/api/v1/user")
	user.Use(mdhw.AuthMiddleware(secret))
	{
		user.GET(":id", userHandler.GetUser)
	}

	// TODO: Probar prometheus
	// Ruta de Prometheus
	r.GET("/metrics", ginInst.WrapH(promhttp.Handler()))

	// Ruta de Salud
	r.GET("/health", userHandler.Health)

	// Integrar Go Micro y Gin
	ms.GetService().Handle("/", r)

	if err := ginInst.RunServer(); err != nil {
		is.MicroLogError("error starting Gin server: %v", err)
	}
}
