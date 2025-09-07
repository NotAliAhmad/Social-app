package api

import (
	"encoding/json"
	"log"
	"net/http"

	"social-app/cmd/internal/models"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// Type aliases for backward compatibility
type CreateUserRequest = models.CreateUserRequest
type UpdateUserRequest = models.UpdateUserRequest
type User = models.User

func (server *Server) listUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := server.Store.ListUsers(r.Context())
	if err != nil {
		http.Error(w, "Failed to list users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (server *Server) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.FirstName == "" || req.LastName == "" {
		http.Error(w, "First name and last name are required", http.StatusBadRequest)
		return
	}

	user := User{
		ID:        uuid.New().String(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	err = server.Store.CreateUser(r.Context(), user)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (server *Server) getUserHandler(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "userID")

	user, err := server.Store.GetUser(r.Context(), ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (server *Server) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "userID")

	var req UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Get the existing user
	existingUser, err := server.Store.GetUser(r.Context(), ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Update the user with new values
	updatedUser := *existingUser
	if req.FirstName != "" {
		updatedUser.FirstName = req.FirstName
	}
	if req.LastName != "" {
		updatedUser.LastName = req.LastName
	}

	err = server.Store.UpdateUser(r.Context(), ID, updatedUser)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(updatedUser)
}

func (server *Server) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "userID")

	err := server.Store.DeleteUser(r.Context(), ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User Deleted"))
}
