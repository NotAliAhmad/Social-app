package api

import (
	"net/http"
	"social-app/cmd/internal/store"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Config Config
	Store  store.Store
}

type Config struct {
	Addr string
}

// NewApplication creates a new Application instance with the provided config and store
func NewServer(config Config, store store.Store) (*Server, error) {
	server := &Server{
		Config: config,
		Store:  store,
	}

	return server, nil
}

func (server *Server) Serve() error {
	srv := &http.Server{
		Addr:         server.Config.Addr,
		Handler:      server.mount(),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	return srv.ListenAndServe()
}

func (server *Server) mount() *chi.Mux {
	chiMux := chi.NewRouter()
	chiMux.Use(middleware.RequestID)
	chiMux.Use(middleware.RealIP)
	chiMux.Use(middleware.Logger)
	chiMux.Use(middleware.Recoverer)

	chiMux.HandleFunc("/v1/healthcheck", server.healthcheckHandler)

	// home page
	chiMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is the home page!"))
	})
	chiMux.HandleFunc("/healthcheck", server.healthcheckHandler)
	// API v1 routes
	//chiMux.Group(func(r chi.Router) {
	//r.Use(server.authMiddleware) // Custom auth middleware

	chiMux.Route("/v1", func(r chi.Router) {
		r.Mount("/auth", server.authRoutes())
		r.Mount("/users", server.userRoutes())
		r.Mount("/posts", server.postRoutes())

	})

	return chiMux
}

func (server *Server) userRoutes() chi.Router {
	r := chi.NewRouter()
	// handlers for user routes

	r.Get("/", server.listUsersHandler)
	r.Post("/", server.createUserHandler)
	r.Route("/{userID}", func(r chi.Router) {
		r.Get("/", server.getUserHandler)
		r.Put("/", server.updateUserHandler)
		r.Delete("/", server.deleteUserHandler)
	})

	return r
}

func (server *Server) postRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", server.listPostsHandler)
	r.Post("/", server.createPostHandler)
	r.Route("/{postID}", func(r chi.Router) {
		r.Get("/", server.getPostHandler)
		r.Put("/", server.updatePostHandler)
		r.Delete("/", server.deletePostHandler)

		// Nested routes
		r.Route("/comments", func(r chi.Router) {
			r.Get("/", server.listPostCommentsHandler)
			r.Post("/", server.createCommentHandler)
		})
	})

	return r
}

func (server *Server) authMiddleware(next http.Handler) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	return mux
}
