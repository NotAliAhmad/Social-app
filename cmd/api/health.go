package api

import "net/http"

func (server *Server) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}