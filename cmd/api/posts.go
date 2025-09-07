package api

import (
	"net/http"
)

func (server *Server) listPostsHandler(w http.ResponseWriter, r *http.Request) {
	// Your logic here
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Posts retrieved"))
}

func (server *Server) createPostHandler(w http.ResponseWriter, r *http.Request) {
	// Your logic here
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Post created"))
}

func (server *Server) getPostHandler(w http.ResponseWriter, r *http.Request) {
	// Your logic here
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post retrieved"))
}

func (server *Server) updatePostHandler(w http.ResponseWriter, r *http.Request) {
	// Your logic here
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post updated"))
}

func (server *Server) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	// Your logic here
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post deleted"))
}

func (server *Server) listPostCommentsHandler(w http.ResponseWriter, r *http.Request) {
	// Your logic here
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post comments retrieved"))
}

func (server *Server) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	// Your logic here
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Comment created"))
}
