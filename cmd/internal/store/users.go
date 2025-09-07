package store

import (
	"context"
	"social-app/cmd/internal/models"
)

type UsersStore interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUser(ctx context.Context, id string) (*models.User, error)
	UpdateUser(ctx context.Context, id string, user models.User) error
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context) ([]models.User, error)
}

// User operations implementation
func (p *PostgresStore) CreateUser(ctx context.Context, user models.User) error {
	query := `INSERT INTO users (id, first_name, last_name) VALUES ($1, $2, $3)`
	_, err := p.db.ExecContext(ctx, query, user.ID, user.FirstName, user.LastName)
	return err
}

func (p *PostgresStore) GetUser(ctx context.Context, id string) (*models.User, error) {
	query := `SELECT id, first_name, last_name FROM users WHERE id = $1`
	row := p.db.QueryRowContext(ctx, query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (p *PostgresStore) UpdateUser(ctx context.Context, id string, user models.User) error {
	query := `UPDATE users SET first_name = $1, last_name = $2 WHERE id = $3`
	_, err := p.db.ExecContext(ctx, query, user.FirstName, user.LastName, id)
	return err
}

func (p *PostgresStore) DeleteUser(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := p.db.ExecContext(ctx, query, id)
	return err
}

func (p *PostgresStore) ListUsers(ctx context.Context) ([]models.User, error) {
	query := `SELECT id, first_name, last_name FROM users ORDER BY first_name`
	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}