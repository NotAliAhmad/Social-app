package store

import (
	"context"
	"database/sql"
	"social-app/cmd/internal/models"
)

type AuthStore interface {
	CreateSession(ctx context.Context, session models.Session) error
	GetSession(ctx context.Context, token string) (*models.Session, error)
	DeleteSession(ctx context.Context, token string) error
}

type AuthPostgresStore struct {
	db *sql.DB
}

func NewAuthPostgresStore(db *sql.DB) *AuthPostgresStore {
	return &AuthPostgresStore{
		db: db,
	}
}

// Session operations (placeholder implementations)
func (a *AuthPostgresStore) CreateSession(ctx context.Context, session models.Session) error {
	// TODO: Implement
	return nil
}

func (a *AuthPostgresStore) GetSession(ctx context.Context, token string) (*models.Session, error) {
	// TODO: Implement
	return nil, nil
}

func (a *AuthPostgresStore) DeleteSession(ctx context.Context, token string) error {
	// TODO: Implement
	return nil
}
