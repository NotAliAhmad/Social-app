package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"social-app/cmd/api"
	"time"

	_ "github.com/lib/pq"
)

// Config holds database configuration
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewConnection creates a new database connection
func NewConnection(dbconfig api.DBConfig) (*sql.DB, error) {
	// Use the connection string from DBConfig.Addr if provided, otherwise use individual env vars
	var dsn string
	if dbconfig.Addr != "" {
		dsn = dbconfig.Addr
	} else {
		config := Config{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "password"),
			DBName:   getEnv("DB_NAME", "social_app"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		}

		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(dbconfig.MaxOpenConns)
	db.SetMaxIdleConns(dbconfig.MaxIdleConns)

	if dbconfig.MaxIdleTime != "" {
		maxtime, err := time.ParseDuration(dbconfig.MaxIdleTime + "m") // assume minutes
		if err != nil {
			return nil, fmt.Errorf("failed to parse MaxIdleTime: %w", err)
		}
		db.SetConnMaxIdleTime(maxtime)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to database") 
	return db, nil
}

// getEnv gets an environment variable with a fallback default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
