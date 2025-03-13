package router

import (
	"net/http"

	httpPkg "blog-api/pkg/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// chiRouter represents the Chi router implementation
type chiRouter struct {
	mux         *chi.Mux
	postHandler httpPkg.PostHandler
}

// NewRouter creates a new router instance
func NewRouter(config httpPkg.RouterConfig) httpPkg.Router {
	r := &chiRouter{
		mux:         chi.NewRouter(),
		postHandler: config.PostHandler,
	}
	r.Setup()
	return r
}

// ServeHTTP implements the http.Handler interface
func (r *chiRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

// Setup sets up all the routes and middleware
func (r *chiRouter) Setup() {
	// Middleware
	r.mux.Use(middleware.Logger)
	r.mux.Use(middleware.Recoverer)

	// Posts routes
	r.mux.Route("/posts", func(router chi.Router) {
		router.Post("/", r.postHandler.CreatePost)
		router.Get("/", r.postHandler.GetAllPosts)
		router.Get("/{id}", r.postHandler.GetPost)
		router.Put("/{id}", r.postHandler.UpdatePost)
	})
}
