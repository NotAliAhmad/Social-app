package main

import (
	"database/sql"
	"log"
	"social-app/cmd/api"
	"social-app/cmd/internal/db"
	"social-app/cmd/internal/env"
	"social-app/cmd/internal/store"
)

func main() {
	cfg := api.Config{
		Addr: env.GetString("ADDR", ":8080"),
		DBconfig: api.DBConfig{
			Addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost:5432/social?sslmode=disable"),
			MaxOpenConns: env.GetInt("DB_MAX_OPEN-CONNS", 30),
			MaxIdleConns: env.GetInt("DB_MAX_OPEN-CONNS", 30),
			MaxIdleTime:  env.GetString("DB_MAX_OPEN-CONNS", "15"),
		},
	}

	// Choose store based on environment
	var str store.Store
	var dbConn *sql.DB
	var err error

	// Use PostgreSQL store
	dbConn, err = db.NewConnection(cfg.DBconfig)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer dbConn.Close()

	str = store.NewPostgresStore(dbConn)
	if err != nil {
		log.Fatal("Failed to create Postgres store:", err)
	}
	log.Println("Using PostgreSQL store")

	// Create the server with the store
	server, err := api.NewServer(cfg, str)
	if err != nil {
		log.Fatal("Failed to create server:", err)
	}

	// // Set the database connection if using PostgreSQL
	// if dbConn != nil {
	// 	server.SetDB(dbConn)
	// }

	err = server.Serve()
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
