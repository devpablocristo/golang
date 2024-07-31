//go:build wireinject
// +build wireinject

package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/google/wire"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"

	hdl "github.com/devpablocristo/qh/events/cmd/api/handlers"
	"github.com/devpablocristo/qh/events/internal/core"
	"github.com/devpablocristo/qh/events/internal/core/event"
	"github.com/devpablocristo/qh/events/internal/platform/config"
	db "github.com/devpablocristo/qh/events/internal/platform/repository"
)

func provideRegistry(config *config.Dependencies) registry.Registry {
	return consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{config.ConsulConfig.Address}
	})
}

func provideMicroservice(reg registry.Registry) web.Service {
	ms := web.NewService(
		web.Name("events"),
		web.Version("latest"),
		web.Registry(reg),
		web.Address(":8888"),
	)

	if err := ms.Init(); err != nil {
		logger.Fatal(err)
	}
	return ms
}

func provideDatabaseInstance(config *config.Dependencies) *db.PostgreSQL {
	return db.NewPostgreSQL(config.DBConfig)
}

func provideRepository(db *db.PostgreSQL) event.RepositoryPort {
	return event.NewRepository(db)
}

func provideUseCase(repo event.RepositoryPort) core.UseCasePort {
	return core.NewUseCase(repo)
}

func provideHandler(useCase core.UseCasePort) *hdl.RestHandler {
	return hdl.NewRestHandler(useCase)
}

func provideRouter(ms web.Service, handler *hdl.RestHandler) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1/events")
	{
		v1.Use(hdl.LoggingMiddleware())
		v1.POST("/fake-create", handler.FakeCreateEvent)
	}

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.GET("/health", handler.Health)
	ms.Handle("/", r)

	return r
}

func InitRouter(config *config.Dependencies) (*gin.Engine, error) {
	wire.Build(
		provideRegistry,
		provideMicroservice,
		provideDatabaseInstance,
		provideRepository,
		provideUseCase,
		provideHandler,
		provideRouter,
	)
	return &gin.Engine{}, nil
}

func registerServiceWithConsul(config *config.Dependencies) error {
	consulConfig := consulapi.DefaultConfig()
	consulConfig.Address = config.ConsulConfig.Address
	client, err := consulapi.NewClient(consulConfig)
	if err != nil {
		return fmt.Errorf("failed to create Consul client: %w", err)
	}

	// Registro del servicio
	registration := &consulapi.AgentServiceRegistration{
		ID:      config.ConsulConfig.ID,      // ID único del servicio
		Name:    config.ConsulConfig.Name,    // Nombre del servicio
		Port:    config.ConsulConfig.Port,    // Puerto del servicio
		Address: config.ConsulConfig.Service, // Dirección del servicio (nombre del contenedor en Docker Compose)
		Check: &consulapi.AgentServiceCheck{
			HTTP:     config.ConsulConfig.HTTP, // URL del health check
			Interval: "10s",                    // Intervalo de chequeo
			Timeout:  "1s",                     // Timeout del chequeo
		},
	}

	return client.Agent().ServiceRegister(registration)
}

func RunServer(router *gin.Engine, config *config.Dependencies) error {
	if err := registerServiceWithConsul(config); err != nil {
		return err
	}
	logger.Info("Registering service with Consul")

	if err := router.Run(":" + config.RouterConfig.RouterPort); err != nil {
		return err
	}
	return nil
}
