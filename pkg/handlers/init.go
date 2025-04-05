package handlers

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/mattn/go-sqlite3"
)

const dbFile = ".todo/database.db"

func InitDB() error {
	dir := ".todo"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0755); err != nil {
			return fmt.Errorf("Failed to create directory %s: %v", dir, err)
		}
	}

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		done BOOLEAN NOT NULL DEFAULT 0
	);`

	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	fmt.Println("✅ Database created successfully")
	return nil
}
