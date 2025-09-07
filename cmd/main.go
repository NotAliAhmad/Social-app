package main

import (
	"log"
	"social-app/cmd/api"
	"social-app/cmd/internal/env"
	"social-app/cmd/internal/store"
)

func main() {
	cfg := api.Config{
		Addr: env.GetString("ADDR", ":8080"),
	}

	// Initialize the store (using memory store for now)
	// You can easily switch to PostgresStore later
	memoryStore := store.NewMemoryStore()
	
	// Create the server with the store
	app, err := api.NewServer(cfg, memoryStore)
	if err != nil {
		log.Fatal("Failed to create server:", err)
	}

	err = app.Serve()
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
