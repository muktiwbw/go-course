package interfaces

import "context"

// Server represents the HTTP server interface
type Server interface {
	Start() error
	Shutdown(ctx context.Context) error
}

// ServerConfig holds the server configuration
type ServerConfig struct {
	Port string
}
