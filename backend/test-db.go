package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Try to load .env
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: .env file not loaded: %v", err)
	}
	
	// Get database config
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres"
	}
	
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		fmt.Print("Enter database password: ")
		fmt.Scanln(&password)
	}
	
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "nftgenie"
	}
	
	// Build connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	
	fmt.Printf("Attempting to connect to database...\n")
	fmt.Printf("Host: %s, Port: %s, User: %s, DB: %s\n", host, port, user, dbname)
	
	// Try to connect
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()
	
	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	
	fmt.Println("✅ Successfully connected to database!")
	
	// List tables
	rows, err := db.Query("SELECT tablename FROM pg_tables WHERE schemaname='public' ORDER BY tablename")
	if err != nil {
		log.Fatalf("Failed to query tables: %v", err)
	}
	defer rows.Close()
	
	fmt.Println("\nTables in database:")
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			continue
		}
		fmt.Printf("  - %s\n", tableName)
	}
}
