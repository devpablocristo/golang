package grpcpkg

import (
	"fmt"

	portspkg "github.com/devpablocristo/golang/sdk/pkg/grpc/google/portspkg"
)

// grpcConfig representa la configuración necesaria para conectarse a un servidor gRPC.
type grpcConfig struct {
	host      string
	port      int
	tlsConfig *portspkg.TLSConfig
}

// NewGrpcConfig crea una nueva configuración gRPC con los valores proporcionados.
func NewGrpcConfig(host string, port int, tlsConfig *portspkg.TLSConfig) portspkg.GrpcConfig {
	return &grpcConfig{
		host:      host,
		port:      port,
		tlsConfig: tlsConfig,
	}
}

// GetHost devuelve el host de la configuración.
func (config *grpcConfig) GetHost() string {
	return config.host
}

// SetHost establece el host de la configuración.
func (config *grpcConfig) SetHost(host string) {
	config.host = host
}

// GetPort devuelve el puerto de la configuración.
func (config *grpcConfig) GetPort() int {
	return config.port
}

// SetPort establece el puerto de la configuración.
func (config *grpcConfig) SetPort(port int) {
	config.port = port
}

// GetTLSConfig devuelve la configuración TLS.
func (config *grpcConfig) GetTLSConfig() *portspkg.TLSConfig {
	return config.tlsConfig
}

// SetTLSConfig establece la configuración TLS.
func (config *grpcConfig) SetTLSConfig(tlsConfig *portspkg.TLSConfig) {
	config.tlsConfig = tlsConfig
}

// Validate valida la configuración gRPC.
func (config *grpcConfig) Validate() error {
	if config.host == "" {
		return fmt.Errorf("gRPC host is not configured")
	}
	if config.port == 0 {
		return fmt.Errorf("gRPC port is not configured")
	}
	return nil
}
