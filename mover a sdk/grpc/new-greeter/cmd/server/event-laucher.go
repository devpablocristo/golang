package main

import (
	"context"
	"errors"

	ctypes "github.com/devpablocristo/qh/internal/platform/custom-types"
	mongodb "github.com/devpablocristo/qh/internal/platform/mongodb"

	handler "github.com/devpablocristo/qh/cmd/server/handlers"
	eveman "github.com/devpablocristo/qh/internal/event-manager"
	event "github.com/devpablocristo/qh/internal/event-manager/event"
)

// LauncherPort is an interface for managing application configuration and services.
type EventLauncherPort interface {
	Setup(ctx context.Context) error
	Stop(ctx context.Context) error
	InitEventService(ctx context.Context) error
}

type EventLauncher struct {
	ginConfig   event.ConfigGinPort
	mongoConfig event.ConfigMongoPort
	mongoDB     event.MongoDBServicePort
	router      handler.PortRouterService
	dao         event.MongoEventDAOPort
	repo        event.RepoPort
	useCase     event.UseCasePort
	handlerr    handler.HandlerPort
	running     bool
}

// NewLauncher creates a new instance of EventLauncher.
func NewLauncher() EventLauncherPort {
	return &EventLauncher{}
}

// Setup initializes application components.
func (l *EventLauncher) Setup(ctx context.Context) error {
	l.setupEventComponents()
	return nil
}

func (l *EventLauncher) InitEventService(ctx context.Context) error {
	if l.running {
		return errors.New("the service is already running")
	}

	if err := l.mongoDB.Connect(ctx); err != nil {
		return ctypes.New("error connecting to MongoDB: %v", err)
	}

	err := l.router.SetupRouter(ctx).Run(":" + l.ginConfig.GetHandlerPort())
	if err != nil {
		l.mongoDB.Disconnect(ctx)
		return ctypes.New("error setting up the routing service", err)
	}

	l.running = true

	return nil
}

// Stop stops the services.
func (l *EventLauncher) Stop(ctx context.Context) error {
	if !l.running {
		return ctypes.New("the events service is not running", nil)
	}

	l.mongoDB.Disconnect(ctx)

	l.running = false
	return nil
}

// setupEventComponents initializes components for the events service.
func (l *EventLauncher) setupEventComponents() {
	l.ginConfig = event.NewGinConfig()
	l.mongoConfig = event.NewMongoConfig()
	l.mongoDB = mongodb.NewMongoDBService(l.mongoConfig)
	l.dao = event.NewMongoEventDAO(l.mongoDB)
	l.repo = event.NewRepo(l.dao)
	l.useCase = eveman.NewUseCase(l.repo)
	l.handlerr = handler.NewHandler(l.useCase)
	l.router = handler.NewRouterService(l.handlerr)
}
