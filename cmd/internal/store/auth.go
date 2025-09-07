package store

import (
	"context"
	"social-app/cmd/internal/models"
)

type authStore interface {
	CreateSession(ctx context.Context, session models.Session) error
	GetSession(ctx context.Context, token string) (*models.Session, error)
	DeleteSession(ctx context.Context, token string) error
}

// Session operations (placeholder implementations)
func (p *PostgresStore) CreateSession(ctx context.Context, session models.Session) error {
	// TODO: Implement
	return nil
}

func (p *PostgresStore) GetSession(ctx context.Context, token string) (*models.Session, error) {
	// TODO: Implement
	return nil, nil
}

func (p *PostgresStore) DeleteSession(ctx context.Context, token string) error {
	// TODO: Implement
	return nil
}
