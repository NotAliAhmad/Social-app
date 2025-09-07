package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (server *Server) authRoutes() chi.Router {
	r := chi.NewRouter()

	// Auth routes
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", server.loginHandler)
		r.Post("/register", server.registerHandler)
		r.Post("/logout", server.logoutHandler)
	})

	return r
}

func (server *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	// Your logic here
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}

func (server *Server) registerHandler(w http.ResponseWriter, r *http.Request) {
	// Your logic here
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Register successful"))
}

func (server *Server) logoutHandler(w http.ResponseWriter, r *http.Request) {
	// Your logic here
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout successful"))
}