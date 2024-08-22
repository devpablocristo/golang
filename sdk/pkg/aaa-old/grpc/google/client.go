package ggrpcgpkg

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/devpablocristo/golang/sdk/pkg/grpc/google/portspkg"
)

var (
	instance  portspkg.GgrpcClient
	once      sync.Once
	initError error
)

type ggrpcClient struct {
	conn *grpc.ClientConn
}

func InitializeGgrpcClient(config portspkg.GgrpcConfig) error {
	once.Do(func() {
		var opts []grpc.DialOption
		if config.GetTLSConfig() != nil {
			tlsConfig, err := loadTLSConfig(config.GetTLSConfig())
			if err != nil {
				initError = fmt.Errorf("failed to load TLS config: %v", err)
				return
			}
			creds := credentials.NewTLS(tlsConfig)
			opts = append(opts, grpc.WithTransportCredentials(creds))
		} else {
			opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
		}

		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.GetHost(), config.GetPort()), opts...)
		if err != nil {
			initError = fmt.Errorf("failed to connect to gRPC server: %v", err)
			return
		}

		instance = &ggrpcClient{conn: conn}
	})
	return initError
}

func GetGgrpcClientInstance() (portspkg.GgrpcClient, error) {
	if instance == nil {
		return nil, fmt.Errorf("grpc client is not initialized")
	}
	return instance, nil
}

// InvokeMethod realiza la invocación de un método gRPC.
func (client *ggrpcClient) InvokeMethod(ctx context.Context, method string, request, response any) error {
	return client.conn.Invoke(ctx, method, request, response)
}

// Close cierra la conexión con el servidor gRPC.
func (client *ggrpcClient) Close() error {
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
