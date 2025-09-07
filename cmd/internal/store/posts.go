package store

import (
	"context"
	"social-app/cmd/internal/models"
)

type PostsStore interface {
	CreatePost(ctx context.Context, post models.Post) error
	GetPost(ctx context.Context, id string) (*models.Post, error)
	UpdatePost(ctx context.Context, id string, post models.Post) error
	DeletePost(ctx context.Context, id string) error
	ListPosts(ctx context.Context) ([]models.Post, error)
}

// Post operations (placeholder implementations)
func (p *PostgresStore) CreatePost(ctx context.Context, post models.Post) error {
	// TODO: Implement
	return nil
}

func (p *PostgresStore) GetPost(ctx context.Context, id string) (*models.Post, error) {
	// TODO: Implement
	return nil, nil
}

func (p *PostgresStore) UpdatePost(ctx context.Context, id string, post models.Post) error {
	// TODO: Implement
	return nil
}

func (p *PostgresStore) DeletePost(ctx context.Context, id string) error {
	// TODO: Implement
	return nil
}

func (p *PostgresStore) ListPosts(ctx context.Context) ([]models.Post, error) {
	// TODO: Implement
	return nil, nil
}