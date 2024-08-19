package ggrpcgpkg

import (
	"fmt"

	portspkg "github.com/devpablocristo/golang/sdk/pkg/grpc/google/portspkg"
)

// ggrpcConfig representa la configuración necesaria para conectarse a un servidor gRPC.
type ggrpcConfig struct {
	host      string
	port      int
	tlsConfig *portspkg.TLSConfig
}

// NewGrpcConfig crea una nueva configuración gRPC con los valores proporcionados.
func NewGrpcConfig(host string, port int, tlsConfig *portspkg.TLSConfig) portspkg.GgrpcConfig {
	return &ggrpcConfig{
		host:      host,
		port:      port,
		tlsConfig: tlsConfig,
	}
}

// GetHost devuelve el host de la configuración.
func (config *ggrpcConfig) GetHost() string {
	return config.host
}

// SetHost establece el host de la configuración.
func (config *ggrpcConfig) SetHost(host string) {
	config.host = host
}

// GetPort devuelve el puerto de la configuración.
func (config *ggrpcConfig) GetPort() int {
	return config.port
}

// SetPort establece el puerto de la configuración.
func (config *ggrpcConfig) SetPort(port int) {
	config.port = port
}

// GetTLSConfig devuelve la configuración TLS.
func (config *ggrpcConfig) GetTLSConfig() *portspkg.TLSConfig {
	return config.tlsConfig
}

// SetTLSConfig establece la configuración TLS.
func (config *ggrpcConfig) SetTLSConfig(tlsConfig *portspkg.TLSConfig) {
	config.tlsConfig = tlsConfig
}

// Validate valida la configuración gRPC.
func (config *ggrpcConfig) Validate() error {
	if config.host == "" {
		return fmt.Errorf("gRPC host is not configured")
	}
	if config.port == 0 {
		return fmt.Errorf("gRPC port is not configured")
	}
	return nil
}
