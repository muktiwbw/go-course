package repository

import (
	"sync"

	"blog-api/internal/domain"
)

type memoryPostRepository struct {
	sync.RWMutex
	posts map[string]*domain.Post
}

// NewMemoryPostRepository creates a new in-memory post repository
func NewMemoryPostRepository() domain.PostRepository {
	return &memoryPostRepository{
		posts: make(map[string]*domain.Post),
	}
}

func (m *memoryPostRepository) Store(post *domain.Post) error {
	m.Lock()
	defer m.Unlock()
	m.posts[post.ID] = post
	return nil
}

func (m *memoryPostRepository) GetByID(id string) (*domain.Post, error) {
	m.RLock()
	defer m.RUnlock()
	if post, exists := m.posts[id]; exists {
		return post, nil
	}
	return nil, domain.ErrPostNotFound
}

func (m *memoryPostRepository) GetAll() ([]*domain.Post, error) {
	m.RLock()
	defer m.RUnlock()
	posts := make([]*domain.Post, 0, len(m.posts))
	for _, post := range m.posts {
		posts = append(posts, post)
	}
	return posts, nil
}

func (m *memoryPostRepository) Update(post *domain.Post) error {
	m.Lock()
	defer m.Unlock()
	if _, exists := m.posts[post.ID]; !exists {
		return domain.ErrPostNotFound
	}
	m.posts[post.ID] = post
	return nil
}
