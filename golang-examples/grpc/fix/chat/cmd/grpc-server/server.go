package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"google.golang.org/grpc"

	grpchandler "chat/internal/adapters/grpc"
	mongodb "chat/internal/adapters/mongodb"
	chat "chat/internal/application"
	config "chat/internal/config"
	pb "chat/internal/pb"
)

func main() {

	cfg := config.LoadConfig()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//logger := loreuslogging.FromContext(ctx)

	//_ = tracer.Setup(ctx)

	db, err := mongodbInit(ctx, cfg)
	if err != nil {
		log.Fatalf("error chat/internal/config:%v", err)
	}
	chatRepo := mongodb.NewRepository(db)

	chatService := chat.NewChatService(chatRepo)
	chatHandler := grpchandler.NewChatHandler(chatService)

	err = grpcInit(ctx, cfg, chatHandler)
	if err != nil {
		log.Fatalln(err)
	}
}

// Initialize MongoDB
func mongodbInit(ctx context.Context, cfg config.Config) (*mongo.Database, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		errCustom := fmt.Errorf("failed to connect to MongoDB: %v", err)
		log.Fatalln(errCustom)
		return nil, errCustom
	}
	defer client.Disconnect(ctx)
	db := client.Database(cfg.MongoDatabase)

	return db, nil
}

// Initialize gRPC server
func grpcInit(ctx context.Context, cfg config.Config, chatHandler *grpchandler.ChatHandler) error {
	lis, err := net.Listen("tcp", cfg.GRPCAddress)
	if err != nil {
		errCustom := fmt.Errorf("failed to listen: %v", err)
		log.Fatalln(errCustom)
		return errCustom
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, chatHandler)

	go func() {
		err = grpcServer.Serve(lis)
	}()
	if err != nil {
		errCustom := fmt.Errorf("failed to serve gRPC: %v", err)
		log.Fatalln(errCustom)
		return errCustom
	}

	// Wait for termination signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	grpcServer.GracefulStop()
	log.Println("Server stopped gracefully")

	return nil
}
