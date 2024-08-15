package grpcpkg

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"sync"

	"github.com/devpablocristo/golang/sdk/pkg/grpc/google/portspkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	instance portspkg.GrpcClient
	once     sync.Once
	errInit  error
)

type grpcClient struct {
	conn *grpc.ClientConn
}

// InitializeGrpcClient inicializa una conexión única a gRPC.
func InitializeGrpcClient(config portspkg.GrpcConfig) error {
	once.Do(func() {
		var opts []grpc.DialOption

		if config.GetTLSConfig() != nil {
			tlsConfig, err := loadTLSConfig(config.GetTLSConfig())
			if err != nil {
				errInit = fmt.Errorf("failed to load TLS config: %v", err)
				return
			}
			creds := credentials.NewTLS(tlsConfig)
			opts = append(opts, grpc.WithTransportCredentials(creds))
		} else {
			opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
		}

		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.GetHost(), config.GetPort()), opts...)
		if err != nil {
			errInit = fmt.Errorf("failed to connect to gRPC server: %v", err)
			return
		}

		instance = &grpcClient{conn: conn}
	})
	return errInit
}

// GetGrpcClientInstance devuelve la instancia del cliente gRPC.
func GetGrpcClientInstance() (portspkg.GrpcClient, error) {
	if instance == nil {
		return nil, fmt.Errorf("grpc client is not initialized")
	}
	return instance, nil
}

// InvokeMethod realiza la invocación de un método gRPC.
func (client *grpcClient) InvokeMethod(ctx context.Context, method string, request, response interface{}) error {
	return client.conn.Invoke(ctx, method, request, response)
}

// Close cierra la conexión con el servidor gRPC.
func (client *grpcClient) Close() error {
	return client.conn.Close()
}

// loadTLSConfig carga la configuración TLS desde los archivos provistos.
func loadTLSConfig(tlsConfig *portspkg.TLSConfig) (*tls.Config, error) {
	certificate, err := tls.LoadX509KeyPair(tlsConfig.CertFile, tlsConfig.KeyFile)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	ca, err := os.ReadFile(tlsConfig.CAFile)
	if err != nil {
		return nil, err
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil, fmt.Errorf("failed to append CA certificates")
	}

	return &tls.Config{
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	}, nil
}
