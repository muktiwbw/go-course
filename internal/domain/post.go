package domain

// Post represents the post entity
type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// PostRepository represents the post's repository contract
type PostRepository interface {
	Store(post *Post) error
	GetByID(id string) (*Post, error)
	GetAll() ([]*Post, error)
	Update(post *Post) error
}

// PostUseCase represents the post's use case contract
type PostUseCase interface {
	CreatePost(post *Post) error
	GetPost(id string) (*Post, error)
	GetAllPosts() ([]*Post, error)
	UpdatePost(post *Post) error
}
