package server

import (
	"context"
	"fmt"
	"net/http"
	"os"

	httpPkg "blog-api/pkg/http"
)

// httpServer represents the HTTP server implementation
type httpServer struct {
	server *http.Server
}

// NewServerConfig creates a new server configuration from environment variables
func NewServerConfig() httpPkg.ServerConfig {
	return httpPkg.ServerConfig{
		Port: os.Getenv("SERVER_PORT"),
	}
}

// NewServer creates a new HTTP server instance
func NewServer(router http.Handler, config httpPkg.ServerConfig) httpPkg.Server {
	return &httpServer{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%s", config.Port),
			Handler: router,
		},
	}
}

// Start starts the HTTP server
func (s *httpServer) Start() error {
	return s.server.ListenAndServe()
}

// Shutdown gracefully shuts down the server
func (s *httpServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
