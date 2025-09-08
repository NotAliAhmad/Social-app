package store

import (
	"context"
	"database/sql"
	"social-app/cmd/internal/models"

	"github.com/lib/pq"
)

type PostStore interface {
	CreatePost(ctx context.Context, post *models.Post) error
	GetPost(ctx context.Context, id string) (*models.Post, error)
	UpdatePost(ctx context.Context, id string, post *models.Post) error
	DeletePost(ctx context.Context, id string) error
	ListPosts(ctx context.Context) ([]models.Post, error)
}

type PostPostgresStore struct {
	db *sql.DB
}

func NewPostPostgresStore(db *sql.DB) *PostPostgresStore {
	return &PostPostgresStore{
		db: db,
	}
}

// CreatePost inserts a new post and returns server-generated fields
func (po *PostPostgresStore) CreatePost(ctx context.Context, post *models.Post) error {
	query := `
		INSERT INTO posts (content, title, user_id, tags)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`
	return po.db.QueryRowContext(ctx, query, post.Content, post.Title, post.UserID, pq.Array(post.Tags)).
		Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
}

func (po *PostPostgresStore) GetPost(ctx context.Context, id string) (*models.Post, error) {
	query := `
		SELECT id, content, title, user_id, tags, created_at, updated_at
		FROM posts
		WHERE id = $1
	`
	row := po.db.QueryRowContext(ctx, query, id)

	var post models.Post
	if err := row.Scan(&post.ID, &post.Content, &post.Title, &post.UserID, pq.Array(&post.Tags), &post.CreatedAt, &post.UpdatedAt); err != nil {
		return nil, err
	}
	return &post, nil
}

func (po *PostPostgresStore) UpdatePost(ctx context.Context, id string, post *models.Post) error {
	query := `
		UPDATE posts
		SET content = $1,
			 title = $2,
			 tags = $3,
			 updated_at = NOW()
		WHERE id = $4
	`
	_, err := po.db.ExecContext(ctx, query, post.Content, post.Title, pq.Array(post.Tags), id)
	return err
}

func (po *PostPostgresStore) DeletePost(ctx context.Context, id string) error {
	query := `DELETE FROM posts WHERE id = $1`
	_, err := po.db.ExecContext(ctx, query, id)
	return err
}

func (po *PostPostgresStore) ListPosts(ctx context.Context) ([]models.Post, error) {
	query := `
		SELECT id, content, title, user_id, tags, created_at, updated_at
		FROM posts
		ORDER BY created_at DESC
	`
	rows, err := po.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(&p.ID, &p.Content, &p.Title, &p.UserID, pq.Array(&p.Tags), &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, rows.Err()
}
