package interfaces

import "net/http"

// PostHandler represents the HTTP handler interface for posts
type PostHandler interface {
	CreatePost(w http.ResponseWriter, r *http.Request)
	GetAllPosts(w http.ResponseWriter, r *http.Request)
	GetPost(w http.ResponseWriter, r *http.Request)
	UpdatePost(w http.ResponseWriter, r *http.Request)
}
