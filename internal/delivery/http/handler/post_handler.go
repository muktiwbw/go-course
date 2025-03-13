package handler

import (
	"encoding/json"
	"net/http"

	"blog-api/internal/domain"
	httpPkg "blog-api/pkg/http"

	"github.com/go-chi/chi/v5"
)

type postHandler struct {
	postUseCase domain.PostUseCase
}

// NewPostHandler creates a new post handler
func NewPostHandler(useCase domain.PostUseCase) httpPkg.PostHandler {
	return &postHandler{
		postUseCase: useCase,
	}
}

func (h *postHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post domain.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.postUseCase.CreatePost(&post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func (h *postHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.postUseCase.GetAllPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func (h *postHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	post, err := h.postUseCase.GetPost(id)
	if err == domain.ErrPostNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func (h *postHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var post domain.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post.ID = id
	if err := h.postUseCase.UpdatePost(&post); err == domain.ErrPostNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
