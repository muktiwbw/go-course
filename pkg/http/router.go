package interfaces

import (
	"net/http"
)

// Router represents the HTTP router interface
type Router interface {
	http.Handler
	Setup()
}

// RouterConfig holds the router configuration
type RouterConfig struct {
	PostHandler PostHandler
}
