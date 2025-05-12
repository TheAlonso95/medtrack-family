package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:postgres@localhost:5432/app_db?sslmode=disable"
)

func main() {
	// Connect to database
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Check connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Println("Connected to database successfully")

	// Seed users
	if err := seedUsers(db); err != nil {
		log.Fatalf("Failed to seed users: %v", err)
	}

	fmt.Println("Database seeded successfully")
}

func seedUsers(db *sql.DB) error {
	// Sample users data
	users := []struct {
		username     string
		email        string
		passwordHash string
		fullName     string
		bio          string
	}{
		{
			username:     "admin",
			email:        "admin@example.com",
			passwordHash: "$2a$10$rNqVDjf6yVr6jLUYXFIQAOi8AhWx/X3QgP5Q6qOQb8jlIKtLf4B3y", // "password"
			fullName:     "Admin User",
			bio:          "System administrator",
		},
		{
			username:     "user1",
			email:        "user1@example.com",
			passwordHash: "$2a$10$rNqVDjf6yVr6jLUYXFIQAOi8AhWx/X3QgP5Q6qOQb8jlIKtLf4B3y", // "password"
			fullName:     "Regular User",
			bio:          "Just a regular user",
		},
	}

	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// Insert users
	for _, user := range users {
		var userID int
		err := tx.QueryRow(`
			INSERT INTO users (username, email, password_hash)
			VALUES ($1, $2, $3)
			RETURNING id
		`, user.username, user.email, user.passwordHash).Scan(&userID)
		if err != nil {
			return fmt.Errorf("failed to insert user %s: %w", user.username, err)
		}

		// Insert profile
		_, err = tx.Exec(`
			INSERT INTO profiles (user_id, full_name, bio)
			VALUES ($1, $2, $3)
		`, userID, user.fullName, user.bio)
		if err != nil {
			return fmt.Errorf("failed to insert profile for user %s: %w", user.username, err)
		}

		fmt.Printf("Seeded user: %s\n", user.username)
	}

	return nil
}