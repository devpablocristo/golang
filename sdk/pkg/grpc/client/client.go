package sdkclient

import (
	"context"
	"fmt"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	ports "github.com/devpablocristo/golang/sdk/pkg/grpc/google/client/ports"
)

var (
	clientInstance ports.Client
	clientOnce     sync.Once
	clientInitErr  error
)

// Client estructura que representa un cliente gRPC
type Client struct {
	conn *grpc.ClientConn
}

// newClient crea una nueva instancia de cliente gRPC
func newClient(config ports.Config) (ports.Client, error) {
	clientOnce.Do(func() {
		var opts []grpc.DialOption
		if config.GetTLSConfig() != nil {
			tlsConfig, err := loadTLSConfig(config.GetTLSConfig())
			if err != nil {
				clientInitErr = fmt.Errorf("failed to load TLS config: %v", err)
				return
			}
			creds := credentials.NewTLS(tlsConfig)
			opts = append(opts, grpc.WithTransportCredentials(creds))
		} else {
			opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
		}

		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.GetHost(), config.GetPort()), opts...)
		if err != nil {
			clientInitErr = fmt.Errorf("failed to connect to gRPC server: %v", err)
			return
		}

		clientInstance = &Client{conn: conn}
	})
	return clientInstance, clientInitErr
}

// Implementación de GetConnection
func (client *Client) GetConnection() (*grpc.ClientConn, error) {
	if client.conn == nil {
		return nil, fmt.Errorf("gRPC client connection is not initialized")
	}
	return client.conn, nil
}

// GetClientInstance devuelve la instancia de cliente gRPC
func GetClientInstance() (ports.Client, error) {
	if clientInstance == nil {
		return nil, fmt.Errorf("gRPC client is not initialized")
	}
	return clientInstance, nil
}

func (client *Client) InvokeMethod(ctx context.Context, method string, request, response any) error {
	return client.conn.Invoke(ctx, method, request, response)
}

func (client *Client) Close() error {
	return client.conn.Close()
}
