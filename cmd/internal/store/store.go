package store

import (
	"database/sql"
)

// Store defines the interface for data persistence operations
type Store interface {
	UsersStore
	PostsStore
	AuthStore
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(db *sql.DB) *PostgresStore {
	return &PostgresStore{
		db: db,
	}
}
