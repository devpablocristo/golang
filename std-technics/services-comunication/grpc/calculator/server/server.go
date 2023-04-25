package main

import (
	"context"
	"log"
	"net"

	"github.com/devpablocristo/golang-examples/grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type calculatorServer struct{}

func (b *calculatorServer) Add(ctx context.Context, req *calculatorpb.AddRequest) (*calculatorpb.AddResponse, error) {

	sum := req.GetNum1() + req.GetNum2()

	res := &calculatorpb.AddResponse{
		Result: sum,
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &calculatorServer{})
	log.Printf("Server listening at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve over port %v: %v", port, err)
	}
}
