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

}
