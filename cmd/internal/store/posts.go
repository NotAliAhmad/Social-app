package store

import (
	"context"
	"database/sql"
	"social-app/cmd/internal/models"
)

type PostsStore interface {
	CreatePost(ctx context.Context, post models.Post) error
	GetPost(ctx context.Context, id string) (*models.Post, error)
	UpdatePost(ctx context.Context, id string, post models.Post) error
	DeletePost(ctx context.Context, id string) error
	ListPosts(ctx context.Context) ([]models.Post, error)
}

type PostsPostgresStore struct {
	db *sql.DB
}

func NewPostsPostgresStore(db *sql.DB) *PostsPostgresStore {
	return &PostsPostgresStore{
		db: db,
	}
}

// Post operations (placeholder implementations)
func (po *PostsPostgresStore) CreatePost(ctx context.Context, post models.Post) error {
	// TODO: Implement
	return nil
}

func (po *PostsPostgresStore) GetPost(ctx context.Context, id string) (*models.Post, error) {
	// TODO: Implement
	return nil, nil
}

func (po *PostsPostgresStore) UpdatePost(ctx context.Context, id string, post models.Post) error {
	// TODO: Implement
	return nil
}

func (po *PostsPostgresStore) DeletePost(ctx context.Context, id string) error {
	// TODO: Implement
	return nil
}

func (po *PostsPostgresStore) ListPosts(ctx context.Context) ([]models.Post, error) {
	// TODO: Implement
	return nil, nil
}
