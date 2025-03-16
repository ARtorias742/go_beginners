package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// global variables to be initialized
var (
	config    AppConfig
	db        *sql.DB
	appLogger *log.Logger
)

// AppConfig holds application configuration
type AppConfig struct {
	AppName    string
	DBPath     string
	LogFile    string
	MaxRetries int
}

// First init: Load configuration
func init() {

	fmt.Println("Initializinng configuration...")
	config = AppConfig{
		AppName:    "AdvancedApp",
		DBPath:     "app.db",
		LogFile:    "app.log",
		MaxRetries: 3,
	}

	// Simulate loading from environment variables (optional override)
	if envDBPath := os.Getenv("DB_PATH"); envDBPath != "" {
		config.DBPath = envDBPath
	}
}

// Second  init: set up logging
func init() {
	fmt.Println("Settign up logging.....")
	file, err := os.OpenFile(config.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	appLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	appLogger.Println("Logger initialized")
}

// Third init: Initialize database connection
func init() {
	fmt.Println("Initializing database connection...")
	var err error
	for i := 0; i < config.MaxRetries; i++ {
		db, err = sql.Open("sqlite3", config.DBPath)
		if err == nil {
			break
		}
		fmt.Printf("DB Connection attempt %d failed: %v\n", i+1, err)
	}

	if err != nil {
		appLogger.Fatalf("Failed to connect to database after %d retries: %v", config.MaxRetries, err)
	}

	// Ensure the database is usable
	err = db.Ping()
	if err != nil {
		appLogger.Fatalf("Database ping failed: %v", err)
	}
	appLogger.Println("Database connection established")
}

// Function tom demonstrate usage
func createTable() {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL
		)`
	_, err := db.Exec(query)
	if err != nil {
		appLogger.Printf("Failed to create table : %v", err)
		return
	}
	appLogger.Println("Users table created or already exists")
}

func main() {
	fmt.Println("\nMain application starting...")
	appLogger.Printf("Starting %s", config.AppName)

	// USe the initialized database
	createTable()

	//Insert a sample record
	_, err := db.Exec("INSERT INTO users (name) VALUES (?)", "Alice")
	if err != nil {
		appLogger.Printf("Failed to insert record: %v", err)
	} else {
		appLogger.Println("Inserted user: Alice")
	}

	fmt.Println("Application running successfully!")
}
