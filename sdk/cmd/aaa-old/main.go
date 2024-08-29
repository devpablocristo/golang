package main

import (
	"log"

	"github.com/devpablocristo/golang/sdk/cmd/gateways/auth"
	msg "github.com/devpablocristo/golang/sdk/cmd/gateways/messaging"
	shared "github.com/devpablocristo/golang/sdk/cmd/gateways/shared"

	ginsetup "github.com/devpablocristo/golang/sdk/internal/bootstrap/gin"
	gmwsetup "github.com/devpablocristo/golang/sdk/internal/bootstrap/go-micro-web"
	inisetup "github.com/devpablocristo/golang/sdk/internal/bootstrap/initial"

	// jwtsetup "github.com/devpablocristo/golang/sdk/internal/bootstrap/jwt"
	amqpsetup "github.com/devpablocristo/golang/sdk/internal/bootstrap/rabbitmq"
)

func main() {
	if err := inisetup.BasicSetup(); err != nil {
		log.Fatalf("Error setting up configurations: %v", err)
	}
	inisetup.LogInfo("Application started with JWT secret key: %s", inisetup.GetJWTSecretKey())
	inisetup.MicroLogInfo("Starting application...")

	// Configurar y verificar Go Micro
	gomicro, err := gmwsetup.NewGoMicroInstance()
	if err != nil {
		inisetup.MicroLogError("error initializing Go Micro: %v", err)
	}

	// jwtpkg, err :=jwtsetup.NewJWTInstance()
	// if err != nil {
	// 	inisetup.MicroLogError("error initializing Gin: %v", err)
	// }

	// Configurar y verificar Gin
	sdkgin, err := ginsetup.NewGinInstance()
	if err != nil {
		inisetup.MicroLogError("error initializing Gin: %v", err)
	}

	r := sdkgin.GetRouter()

	// monitoringHandler, err := shared.InitializeMonitoring()
	// if err != nil {
	// 	inisetup.MicroLogError("userHandler error: %v", err)
	// }
	// monitoring.Routes(sdkgin, monitoringHandler)

	// userHandler, err := shared.InitializeUserHandler()
	// if err != nil {
	// 	inisetup.MicroLogError("userHandler error: %v", err)
	// }

	authHandler, err := shared.InitializeAuthHandler()
	if err != nil {
		inisetup.MicroLogError("userHandler error: %v", err)
	}

	apiVersion := "/api/v1"
	auth.Routes(r, authHandler, apiVersion, "secret")
	//user.Routes(r, userHandler)
	// Manejar rutas HTTP
	gomicro.GetWebService().Handle("/", r)

	// Usar gRPC cliente o servidor
	grpcClient := gomicro.GetGrpcClient()
	grpcServer := gomicro.GetGrpcServer()

	_ = grpcClient
	_ = grpcServer
	// Ejecuta Gin en la dirección especificada por Go-Micro
	go func() {
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("Failed to run Gin: %v", err)
		}
	}()

	// Iniciar mensajería (productor y consumidor)
	go messaging()

	// Ejecutar el servicio Go Micro
	if err := gomicro.GetWebService().Run(); err != nil {
		inisetup.MicroLogError("error starting GoMicro service: %v", err)
	}
}

// RabbitMQ
func messaging() {
	client, err := amqpsetup.NewRabbitMqInstance()
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ client: %v", err)
	}

	c, err := client.Channel()
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ chan: %v", err)
	}

	// Iniciar consumidor
	go msg.StartConsumer(c, "exampleQueue")

	// Iniciar productor
	go msg.StartProducer(c, "exampleQueue")
}

//client chat grpc

// package main

// import (
// 	"bufio"
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	pb "chat/internal/pb"

// 	"google.golang.org/grpc"
// )

// func main() {
// 	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("Failed to connect to server: %v", err)
// 	}
// 	defer conn.Close()

// 	client := pb.NewChatServiceClient(conn)

// 	stream, err := client.Chat(context.Background())
// 	if err != nil {
// 		log.Fatalf("Error opening chat stream: %v", err)
// 	}

// 	// Start a goroutine to listen for server messages
// 	go func() {
// 		for {
// 			msg, err := stream.Recv()
// 			if err != nil {
// 				log.Fatalf("Error receiving message from server: %v", err)
// 			}
// 			fmt.Printf("%s: %s\n", msg.Sender, msg.Message)
// 		}
// 	}()

// 	// Main loop to send messages to the server
// 	scanner := bufio.NewScanner(os.Stdin)
// 	for {
// 		fmt.Print("Enter your message (or 'exit' to quit): ")
// 		scanner.Scan()
// 		text := scanner.Text()

// 		if text == "exit" {
// 			break
// 		}

// 		msg := &pb.ChatMessage{
// 			Sender:    "Client",
// 			Message:   text,
// 			Timestamp: time.Now().Unix(),
// 		}

// 		if err := stream.Send(msg); err != nil {
// 			log.Fatalf("Failed to send message to server: %v", err)
// 		}
// 	}
// }

// server char grpc

// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net"
// 	"os"
// 	"os/signal"
// 	"syscall"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"

// 	"google.golang.org/grpc"

// 	grpchandler "chat/internal/adapters/grpc"
// 	mongodb "chat/internal/adapters/mongodb"
// 	chat "chat/internal/application"
// 	config "chat/internal/config"
// 	pb "chat/internal/pb"
// )

// func main() {

// 	cfg := config.LoadConfig()
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	//logger := loreuslogging.FromContext(ctx)

// 	//_ = tracer.Setup(ctx)

// 	db, err := mongodbInit(ctx, cfg)
// 	if err != nil {
// 		log.Fatalf("error chat/internal/config:%v", err)
// 	}
// 	chatRepo := mongodb.NewRepository(db)

// 	chatService := chat.NewChatService(chatRepo)
// 	chatHandler := grpchandler.NewChatHandler(chatService)

// 	err = grpcInit(ctx, cfg, chatHandler)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

// // Initialize MongoDB
// func mongodbInit(ctx context.Context, cfg config.Config) (*mongo.Database, error) {
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
// 	if err != nil {
// 		errCustom := fmt.Errorf("failed to connect to MongoDB: %v", err)
// 		log.Fatalln(errCustom)
// 		return nil, errCustom
// 	}
// 	defer client.Disconnect(ctx)
// 	db := client.Database(cfg.MongoDatabase)

// 	return db, nil
// }

// // Initialize gRPC server
// func grpcInit(ctx context.Context, cfg config.Config, chatHandler *grpchandler.ChatHandler) error {
// 	lis, err := net.Listen("tcp", cfg.GRPCAddress)
// 	if err != nil {
// 		errCustom := fmt.Errorf("failed to listen: %v", err)
// 		log.Fatalln(errCustom)
// 		return errCustom
// 	}

// 	grpcServer := grpc.NewServer()
// 	pb.RegisterChatServiceServer(grpcServer, chatHandler)

// 	go func() {
// 		err = grpcServer.Serve(lis)
// 	}()
// 	if err != nil {
// 		errCustom := fmt.Errorf("failed to serve gRPC: %v", err)
// 		log.Fatalln(errCustom)
// 		return errCustom
// 	}

// 	// Wait for termination signal
// 	sigCh := make(chan os.Signal, 1)
// 	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
// 	<-sigCh

// 	grpcServer.GracefulStop()
// 	log.Println("Server stopped gracefully")

// 	return nil
//}




// client grpc calculator
// func main() {

// 	fmt.Println("Calculator Client")
// 	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("could not connect: %v", err)
// 	}
// 	defer cc.Close()

// 	c := pb.NewCalculatorServiceClient(cc)
// 	// fmt.Printf("Created client: %f", c)

// 	// doUnary(c)

// 	// doServerStreaming(c)

// 	// doClientStreaming(c)

// 	// doBiDiStreaming(c)

// 	doErrorUnary(c)
// }

// func doUnary(c pb.CalculatorServiceClient) {
// 	fmt.Println("Starting to do a Sum Unary RPC...")
// 	req := &pb.SumRequest{
// 		FirstNumber:  5,
// 		SecondNumber: 40,
// 	}
// 	res, err := c.Sum(context.Background(), req)
// 	if err != nil {
// 		log.Fatalf("error while calling Sum RPC: %v", err)
// 	}
// 	log.Printf("Response from Sum: %v", res.SumResult)
// }

// func doServerStreaming(c pb.CalculatorServiceClient) {
// 	fmt.Println("Starting to do a PrimeDecomposition Server Streaming RPC...")
// 	req := &pb.PrimeNumberDecompositionRequest{
// 		Number: 12390392840,
// 	}
// 	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
// 	if err != nil {
// 		log.Fatalf("error while calling PrimeDecomposition RPC: %v", err)
// 	}
// 	for {
// 		res, err := stream.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatalf("Something happened: %v", err)
// 		}
// 		fmt.Println(res.GetPrimeFactor())
// 	}
// }

// func doClientStreaming(c pb.CalculatorServiceClient) {
// 	fmt.Println("Starting to do a ComputeAverage Client Streaming RPC...")

// 	stream, err := c.ComputeAverage(context.Background())
// 	if err != nil {
// 		log.Fatalf("Error while opening stream: %v", err)
// 	}

// 	numbers := []int32{3, 5, 9, 54, 23}

// 	for _, number := range numbers {
// 		fmt.Printf("Sending number: %v\n", number)
// 		stream.Send(&pb.ComputeAverageRequest{
// 			Number: number,
// 		})
// 	}

// 	res, err := stream.CloseAndRecv()
// 	if err != nil {
// 		log.Fatalf("Error while receiving response: %v", err)
// 	}

// 	fmt.Printf("The Average is: %v\n", res.GetAverage())
// }

// func doBiDiStreaming(c pb.CalculatorServiceClient) {
// 	fmt.Println("Starting to do a FindMaximum BiDi Streaming RPC...")

// 	stream, err := c.FindMaximum(context.Background())

// 	if err != nil {
// 		log.Fatalf("Error while opening stream and calling FindMaximum: %v", err)
// 	}

// 	waitc := make(chan struct{})

// 	// send go routine
// 	go func() {
// 		numbers := []int32{4, 7, 2, 19, 4, 6, 32}
// 		for _, number := range numbers {
// 			fmt.Printf("Sending number: %v\n", number)
// 			stream.Send(&pb.FindMaximumRequest{
// 				Number: number,
// 			})
// 			time.Sleep(1000 * time.Millisecond)
// 		}
// 		stream.CloseSend()
// 	}()
// 	// receive go routine
// 	go func() {
// 		for {
// 			res, err := stream.Recv()
// 			if err == io.EOF {
// 				break
// 			}
// 			if err != nil {
// 				log.Fatalf("Problem while reading server stream: %v", err)
// 				break
// 			}
// 			maximum := res.GetMaximum()
// 			fmt.Printf("Received a new maximum of...: %v\n", maximum)
// 		}
// 		close(waitc)
// 	}()
// 	<-waitc
// }

// func doErrorUnary(c pb.CalculatorServiceClient) {
// 	fmt.Println("Starting to do a SquareRoot Unary RPC...")

// 	// correct call
// 	doErrorCall(c, 10)

// 	// error call
// 	doErrorCall(c, -2)
// }

// func doErrorCall(c pb.CalculatorServiceClient, n int32) {
// 	res, err := c.SquareRoot(context.Background(), &pb.SquareRootRequest{Number: n})

// 	if err != nil {
// 		respErr, ok := status.FromError(err)
// 		if ok {
// 			// actual error from gRPC (user error)
// 			fmt.Printf("Error message from server: %v\n", respErr.Message())
// 			fmt.Println(respErr.Code())
// 			if respErr.Code() == codes.InvalidArgument {
// 				fmt.Println("We probably sent a negative number!")
// 				return
// 			}
// 		} else {
// 			log.Fatalf("Big Error calling SquareRoot: %v", err)
// 			return
// 		}
// 	}
// 	fmt.Printf("Result of square root of %v: %v\n", n, res.GetNumberRoot())
// }

// server grpc calculator

// type server struct {
// 	pb.CalculatorServiceServer
// }

// func (*server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
// 	fmt.Printf("Received Sum RPC: %v\n", req)
// 	firstNumber := req.FirstNumber
// 	secondNumber := req.SecondNumber
// 	sum := firstNumber + secondNumber
// 	res := &pb.SumResponse{
// 		SumResult: sum,
// 	}
// 	return res, nil
// }

// func (*server) PrimeNumberDecomposition(req *pb.PrimeNumberDecompositionRequest, stream pb.CalculatorService_PrimeNumberDecompositionServer) error {
// 	fmt.Printf("Received PrimeNumberDecomposition RPC: %v\n", req)

// 	number := req.GetNumber()
// 	divisor := int64(2)

// 	for number > 1 {
// 		if number%divisor == 0 {
// 			stream.Send(&pb.PrimeNumberDecompositionResponse{
// 				PrimeFactor: divisor,
// 			})
// 			number = number / divisor
// 		} else {
// 			divisor++
// 			fmt.Printf("Divisor has increased to %v\n", divisor)
// 		}
// 	}
// 	return nil
// }

// func (*server) ComputeAverage(stream pb.CalculatorService_ComputeAverageServer) error {
// 	fmt.Printf("Received ComputeAverage RPC\n")

// 	sum := int32(0)
// 	count := 0

// 	for {
// 		req, err := stream.Recv()
// 		if err == io.EOF {
// 			average := float64(sum) / float64(count)
// 			return stream.SendAndClose(&pb.ComputeAverageResponse{
// 				Average: average,
// 			})
// 		}
// 		if err != nil {
// 			log.Fatalf("Error while reading client stream: %v", err)
// 		}
// 		sum += req.GetNumber()
// 		count++
// 	}

// }

// func (*server) FindMaximum(stream pb.CalculatorService_FindMaximumServer) error {
// 	fmt.Println("Received FindMaximum RPC")
// 	maximum := int32(0)

// 	for {
// 		req, err := stream.Recv()
// 		if err == io.EOF {
// 			return nil
// 		}
// 		if err != nil {
// 			log.Fatalf("Error while reading client stream: %v", err)
// 			return err
// 		}
// 		number := req.GetNumber()
// 		if number > maximum {
// 			maximum = number
// 			sendErr := stream.Send(&pb.FindMaximumResponse{
// 				Maximum: maximum,
// 			})
// 			if sendErr != nil {
// 				log.Fatalf("Error while sending data to client: %v", sendErr)
// 				return sendErr
// 			}
// 		}
// 	}
// }

// func (*server) SquareRoot(ctx context.Context, req *pb.SquareRootRequest) (*pb.SquareRootResponse, error) {
// 	fmt.Println("Received SquareRoot RPC")
// 	number := req.GetNumber()
// 	if number < 0 {
// 		return nil, status.Errorf(
// 			codes.InvalidArgument,
// 			fmt.Sprintf("Received a negative number: %v", number),
// 		)
// 	}
// 	return &pb.SquareRootResponse{
// 		NumberRoot: math.Sqrt(float64(number)),
// 	}, nil
// }

// func main() {
// 	fmt.Println("Calculator Server")

// 	lis, err := net.Listen("tcp", "0.0.0.0:50051")
// 	if err != nil {
// 		log.Fatalf("Failed to listen: %v", err)
// 	}

// 	s := grpc.NewServer()
// 	pb.RegisterCalculatorServiceServer(s, &server{})

// 	// Register reflection service on gRPC server.
// 	reflection.Register(s)

// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }

// arithmetic calc main grpc
// func main() {
// 	dbaseDriver := os.Getenv("DB_DRIVER")
// 	dsourceName := os.Getenv("DS_NAME")
// 	var err error
// 	//ports
// 	var dbaseAdaptor ports.DbPort
// 	var core ports.ArithmaticPort
// 	var appAdaptor ports.APIPort
// 	var gRPCAdaptor ports.GRPCPort

// 	//dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)

// 	dbaseAdaptor, err = db.NewAdaptor(dbaseDriver, dsourceName)
// 	if err != nil {
// 		log.Fatal("failed to initiate dbase connection: %v", err)
// 	}
// 	defer dbaseAdaptor.CloseDbConnection()

// 	core = arithamtic.NewAdaptor()
// 	appAdaptor = api.NewApplication(dbaseAdaptor, core)
// 	gRPCAdaptor = gRPC.NewAdaptor(appAdaptor)
// 	gRPCAdaptor.Run()

// 	// creation of type adaptor which has access to all methods
// 	//arithAdaptor := arithmatic.NewAdaptor()

// }

// client gripc bookstore
// package main

// import (
// 	"context"
// 	"log"
// 	"time"

// 	pb "github.com/devpablocristo/golang-examples/grpc/bookstore/book"
// 	"google.golang.org/grpc"
// )

// const (
// 	port = ":9111"
// )

// func main() {
// 	conn, err := grpc.Dial(port, grpc.WithInsecure(), grpc.WithBlock())
// 	if err != nil {
// 		log.Fatalf("Connection failed: %v", err)
// 	}
// 	defer conn.Close()
// 	c := pb.NewBookstoreInventoryClient(conn)

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()

// 	var newBooks = make(map[string]string)

// 	newBooks["Book A"] = "Author A"
// 	newBooks["Book B"] = "Author B"

// 	for title, author := range newBooks {
// 		r, err := c.CreateNewBook(ctx, &pb.NewBook{Title: title, Author: author})
// 		if err != nil {
// 			log.Fatalf("Could not create book: %v", err)
// 		}
// 		log.Printf(`Book details:
// 		Title: %s
// 		Author: %s
// 		Id: %d`, r.GetTitle(), r.GetAuthor(), r.GetId())
// 	}
// }

// server grpc bookstore
// package main

// import (
// 	"context"
// 	"log"
// 	"math/rand"
// 	"net"

// 	pb "github.com/devpablocristo/golang-examples/grpc/bookstore/book"
// 	"google.golang.org/grpc"
// )

// const (
// 	port = ":9111"
// )

// type BookServer struct {
// 	pb.UnimplementedBookstoreInventoryServer
// }

// func (b *BookServer) CreateNewBook(ctx context.Context, nb *pb.NewBook) (*pb.Book, error) {
// 	log.Printf("Recibed: %v", nb.GetTitle())

// 	book := pb.Book{
// 		Title:  nb.GetTitle(),
// 		Author: nb.GetAuthor(),
// 		Id:     rand.Int31n(1000),
// 	}

// 	return &book, nil
// }

// func main() {
// 	lis, err := net.Listen("tcp", port)
// 	if err != nil {
// 		log.Fatalf("Failed to listen on port %s: %v", port, err)
// 	}

// 	s := grpc.NewServer()
// 	pb.RegisterBookstoreInventoryServer(s, &BookServer{})
// 	log.Printf("Server listening at %v", lis.Addr())
// 	err = s.Serve(lis)
// 	if err != nil {
// 		log.Fatalf("Failed to serve over port %v: %v", port, err)
// 	}
// }
