package store

import (
	"database/sql"
)

// Store defines the interface for data persistence operations
type Store interface {
	UserStore
	PostStore
	AuthStore
}

type PostgresStore struct {
	*UserPostgresStore
	*PostPostgresStore
	*AuthPostgresStore
}

func NewPostgresStore(db *sql.DB) *PostgresStore {
	return &PostgresStore{
		UserPostgresStore: NewUsersPostgresStore(db),
		PostPostgresStore: NewPostPostgresStore(db),
		AuthPostgresStore: NewAuthPostgresStore(db),
	}
}
