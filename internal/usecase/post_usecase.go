package usecase

import "blog-api/internal/domain"

type postUseCase struct {
	postRepo domain.PostRepository
}

// NewPostUseCase creates a new post use case
func NewPostUseCase(repo domain.PostRepository) domain.PostUseCase {
	return &postUseCase{
		postRepo: repo,
	}
}

func (p *postUseCase) CreatePost(post *domain.Post) error {
	return p.postRepo.Store(post)
}

func (p *postUseCase) GetPost(id string) (*domain.Post, error) {
	return p.postRepo.GetByID(id)
}

func (p *postUseCase) GetAllPosts() ([]*domain.Post, error) {
	return p.postRepo.GetAll()
}

func (p *postUseCase) UpdatePost(post *domain.Post) error {
	return p.postRepo.Update(post)
}
