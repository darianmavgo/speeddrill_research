package main

import (
	"log"
	"os"

	"github.com/ncruces/go-sqlite3/driver"
	"github.com/ncruces/go-sqlite3/vfs"
)

func main() {
	// Create a new WASM SQLite database file.
	dbFile, err := os.CreateTemp("", "wasm-sqlite-example.db")
	if err != nil {
		log.Fatalf("Failed to create temporary database file: %v", err)
	}
	defer dbFile.Close()

	// Create a new SQLite VFS for the database file.
	vfs := vfs.NewMemDb()

	// Open the SQLite database.
	db, err := driver.Open("sqlite3", vfs, dbFile.Name(), db.DefaultOptions())
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Create a table in the database.
	_, err = db.Exec(`CREATE TABLE users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL
    )`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	// Insert a row into the database.
	stmt, err := db.Prepare(`INSERT INTO users (name) VALUES (?)`)
	if err != nil {
		log.Fatalf("Failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("Alice")
	if err != nil {
		log.Fatalf("Failed to insert row: %v", err)
	}

	// Select all rows from the database.
	rows, err := db.Query(`SELECT * FROM users`)
	if err != nil {
		log.Fatalf("Failed to query database: %v", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		users = append(users, user)
	}

	// Print the results.
	for _, user := range users {
		log.Printf("User: %d - %s", user.ID, user.Name)
	}
}

type User struct {
	ID   int
	Name string
}
