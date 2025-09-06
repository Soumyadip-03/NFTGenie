package database

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

// Config holds database configuration
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// Initialize sets up the database connection
func Initialize() error {
	config := Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", ""),
		DBName:   getEnv("DB_NAME", "nftgenie"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	var err error
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Configure connection pool
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)

	// Test the connection
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connection established successfully")
	return nil
}

// Close closes the database connection
func Close() {
	if DB != nil {
		DB.Close()
	}
}

// Transaction helper for executing functions within a transaction
func Transaction(fn func(*sqlx.Tx) error) error {
	tx, err := DB.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// getEnv gets environment variable with fallback
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// Migrate runs the database migrations
func Migrate() error {
	schemaPath := "database/schema.sql"
	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %w", err)
	}

	_, err = DB.Exec(string(schema))
	if err != nil {
		// Check if it's a "already exists" error - this is okay
		if containsAny(err.Error(), []string{"already exists", "duplicate key"}) {
			log.Println("Database schema already up to date")
			return nil
		}
		return fmt.Errorf("failed to execute schema: %w", err)
	}

	log.Println("Database migration completed successfully")
	return nil
}

// containsAny checks if the string contains any of the substrings
func containsAny(str string, substrings []string) bool {
	for _, substr := range substrings {
		if len(substr) > 0 && len(str) >= len(substr) && 
		   (str == substr || 
		    (len(str) > len(substr) && 
		     (str[:len(substr)] == substr || 
		      str[len(str)-len(substr):] == substr || 
		      findSubstring(str, substr)))) {
			return true
		}
	}
	return false
}

// findSubstring checks if substring exists in string
func findSubstring(str, substr string) bool {
	for i := 0; i <= len(str)-len(substr); i++ {
		if str[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// Helper functions for common queries

// Exists checks if a record exists
func Exists(query string, args ...interface{}) (bool, error) {
	var exists bool
	query = fmt.Sprintf("SELECT EXISTS(%s)", query)
	err := DB.Get(&exists, query, args...)
	return exists, err
}

// Count returns the count of records
func Count(query string, args ...interface{}) (int, error) {
	var count int
	err := DB.Get(&count, query, args...)
	return count, err
}
