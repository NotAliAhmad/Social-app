package store

import (
	"context"
	"sync"

	"social-app/cmd/internal/models"
)

// MemoryStore is an in-memory implementation of the Store interface
type MemoryStore struct {
	users    []models.User
	posts    []models.Post
	sessions []models.Session
	mu       sync.RWMutex
}

// NewMemoryStore creates a new in-memory store
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		users:    make([]models.User, 0),
		posts:    make([]models.Post, 0),
		sessions: make([]models.Session, 0),
	}
}

// User operations
func (m *MemoryStore) CreateUser(ctx context.Context, user models.User) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.users = append(m.users, user)
	return nil
}

func (m *MemoryStore) GetUser(ctx context.Context, id string) (*models.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, user := range m.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, ErrNotFound
}

func (m *MemoryStore) UpdateUser(ctx context.Context, id string, updatedUser models.User) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, user := range m.users {
		if user.ID == id {
			m.users[i] = updatedUser
			return nil
		}
	}
	return ErrNotFound
}

func (m *MemoryStore) DeleteUser(ctx context.Context, id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, user := range m.users {
		if user.ID == id {
			m.users = append(m.users[:i], m.users[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

func (m *MemoryStore) ListUsers(ctx context.Context) ([]models.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	// Return a copy to prevent external modifications
	users := make([]models.User, len(m.users))
	copy(users, m.users)
	return users, nil
}

// Post operations (placeholder implementations)
func (m *MemoryStore) CreatePost(ctx context.Context, post models.Post) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.posts = append(m.posts, post)
	return nil
}

func (m *MemoryStore) GetPost(ctx context.Context, id string) (*models.Post, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, post := range m.posts {
		if post.ID == id {
			return &post, nil
		}
	}
	return nil, ErrNotFound
}

func (m *MemoryStore) UpdatePost(ctx context.Context, id string, updatedPost models.Post) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, post := range m.posts {
		if post.ID == id {
			m.posts[i] = updatedPost
			return nil
		}
	}
	return ErrNotFound
}

func (m *MemoryStore) DeletePost(ctx context.Context, id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, post := range m.posts {
		if post.ID == id {
			m.posts = append(m.posts[:i], m.posts[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

func (m *MemoryStore) ListPosts(ctx context.Context) ([]models.Post, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	posts := make([]models.Post, len(m.posts))
	copy(posts, m.posts)
	return posts, nil
}

// Session operations (placeholder implementations)
func (m *MemoryStore) CreateSession(ctx context.Context, session models.Session) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.sessions = append(m.sessions, session)
	return nil
}

func (m *MemoryStore) GetSession(ctx context.Context, token string) (*models.Session, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, session := range m.sessions {
		if session.Token == token {
			return &session, nil
		}
	}
	return nil, ErrNotFound
}

func (m *MemoryStore) DeleteSession(ctx context.Context, token string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, session := range m.sessions {
		if session.Token == token {
			m.sessions = append(m.sessions[:i], m.sessions[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}
