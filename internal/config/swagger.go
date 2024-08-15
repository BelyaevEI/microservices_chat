package config

import (
	"net"
	"os"

	"github.com/pkg/errors"
)

const (
	swaggerHostEnvName = "SWAGGER_HOST"
	swaggerPortEnvName = "SWAGGER_PORT"
)

// SwaggerConfig config for Swagger
type SwaggerConfig interface {
	Address() string
}

type swaggerConfig struct {
	host string
	port string
}

// NewSwaggerConfig initializes a Swagger configuration.
func NewSwaggerConfig() (SwaggerConfig, error) {
	host := os.Getenv(swaggerHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("swagger host not found")
	}

	port := os.Getenv(swaggerPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("swagger port not found")
	}

	return &swaggerConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *swaggerConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
