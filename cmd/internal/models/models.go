package models

// User represents a user in the system
type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}

// Post represents a social media post
type Post struct {
	ID        string   `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	UserID    string   `json:"userId"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
}

// Session represents a user session
type Session struct {
	Token     string `json:"token"`
	UserID    string `json:"userId"`
	ExpiresAt string `json:"expiresAt"`
}

// CreateUserRequest represents the request payload for creating a user
type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UpdateUserRequest represents the request payload for updating a user
type UpdateUserRequest struct {
	Username *string `json:"username,omitempty"`
	Email    *string `json:"email,omitempty"`
}
