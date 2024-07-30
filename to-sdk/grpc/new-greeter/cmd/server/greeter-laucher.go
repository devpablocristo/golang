package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	ctypes "github.com/devpablocristo/qh/internal/platform/custom-types"
	mongodb "github.com/devpablocristo/qh/internal/platform/mongodb"

	gtrhandler "github.com/devpablocristo/qh/internal/greeter/adapters/handler"
	grtrepo "github.com/devpablocristo/qh/internal/greeter/adapters/persistence"
	gtrconfig "github.com/devpablocristo/qh/internal/greeter/config"
	gtrport "github.com/devpablocristo/qh/internal/greeter/ports"
	gtrpb "github.com/devpablocristo/qh/internal/greeter/proto"
	gtrservice "github.com/devpablocristo/qh/internal/greeter/service"
)

// AppLauncherPort is an interface for managing application configuration and services.
type GreeterLauncherPort interface {
	Setup(ctx context.Context) error
	Stop(ctx context.Context) error
	InitGreeterServer(ctx context.Context) error
}

type GreeterLaucher struct {
	grGrpcConfig  ctypes.ConfigGrpcPort
	grMongoConfig ctypes.ConfigMongoPort
	grMongoDB     grtrepo.MongoDBServicePort
	grDAO         grtrepo.PortMongoGreetDAO
	grRepo        gtrport.Repo
	grService     gtrport.Service
	grHandler     gtrpb.GreeterServer
	grRunning     bool
}

// NewAppLauncher creates a new instance of GreeterLaucher.
func NewAppLauncher() GreeterLauncherPort {
	return &GreeterLaucher{}
}

// Setup initializes application components.
func (l *GreeterLaucher) Setup(ctx context.Context) error {
	l.setupGreeterComponents()
	return nil
}

// Stop stops the services.
func (l *GreeterLaucher) Stop(ctx context.Context) error {
	if !l.grRunning {
		return ctypes.New("the Greeter service is not running", nil)
	}

	l.grMongoDB.Disconnect(ctx)

	l.grRunning = false

	return nil
}

// setupGreeterComponents initializes components for the Greeter service.
func (l *GreeterLaucher) setupGreeterComponents() {
	l.grGrpcConfig = gtrconfig.NewGrpcConfig()
	l.grMongoConfig = gtrconfig.NewMongoConfig()
	l.grMongoDB = mongodb.NewMongoDBService(l.grMongoConfig)
	l.grDAO = grtrepo.NewMongoGreetDAO(l.grMongoDB)
	l.grRepo = grtrepo.NewRepo(l.grDAO)
	l.grService = gtrservice.NewGreetService(l.grRepo)
	l.grHandler = gtrhandler.NewGreetHandler(l.grService)
}

func (l *GreeterLaucher) InitGreeterServer(ctx context.Context) error {
	lis, err := net.Listen("tcp", l.grGrpcConfig.GetServerHost()+":"+l.grGrpcConfig.GetServerPort())
	if err != nil {
		customErr := fmt.Errorf("error starting the Greeter server: %v", err)
		ctypes.HandleFatalError("error starting the Greeter gRPC service", customErr)
		return customErr
	}

	grpcServer := grpc.NewServer()

	gtrpb.RegisterGreeterServer(grpcServer, l.grHandler)

	log.Printf("Greeter gRPC server started at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		customErr := fmt.Errorf("error starting the Greeter gRPC server: %v", err)
		ctypes.HandleFatalError("error starting the Greeter gRPC service", customErr)
		return customErr
	}

	l.grRunning = true
	return nil
}
