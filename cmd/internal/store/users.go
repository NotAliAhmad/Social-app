package store

import (
	"context"
	"database/sql"
	"social-app/cmd/internal/models"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, id string) (*models.User, error)
	UpdateUser(ctx context.Context, id string, user *models.User) error
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context) ([]models.User, error)
}

type UserPostgresStore struct {
	db *sql.DB
}

func NewUsersPostgresStore(db *sql.DB) *UserPostgresStore {
	return &UserPostgresStore{
		db: db,
	}
}

// User operations implementation
func (u *UserPostgresStore) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (id, username, password, email) VALUES ($1, $2, $3, $4) RETURNING created_at`
	err := u.db.QueryRowContext(ctx, query, user.ID, user.Username, user.Password, user.Email).Scan(&user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserPostgresStore) GetUser(ctx context.Context, id string) (*models.User, error) {
	query := `SELECT id, username, password, email, created_at FROM users WHERE id = $1`
	row := u.db.QueryRowContext(ctx, query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserPostgresStore) UpdateUser(ctx context.Context, id string, user *models.User) error {
	query := `UPDATE users SET username = $1, email = $2 WHERE id = $3`
	_, err := u.db.ExecContext(ctx, query, user.Username, user.Email, id)
	return err
}

func (u *UserPostgresStore) DeleteUser(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := u.db.ExecContext(ctx, query, id)
	return err
}

func (u *UserPostgresStore) ListUsers(ctx context.Context) ([]models.User, error) {
	query := `SELECT id, username, password, email, created_at FROM users ORDER BY username`
	rows, err := u.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}
