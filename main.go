package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	_ "github.com/mattn/go-sqlite3"
)

const dbFile = ".todo/database.db"

func initDB() error {
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



func main()  {

	if len(os.Args) < 2 {
		fmt.Println("❌ You should type a command\nUse: todo init")
		return
	}

	command := strings.ToLower(os.Args[1])

	switch command {
	case "init":
		if err := initDB(); err != nil {
			log.Fatal("Error to initialized a database:\n", err)
		}
	case "start":
		db, err := sql.Open("sqlite3", dbFile)
		if err != nil {
			log.Fatal("Failed to open database:", err)
		}
		defer db.Close()

		query := `SELECT id, title, done FROM tasks;`

		rows, err := db.Query(query)

		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		for rows.Next() {
			var id int
			var title string
			var done bool

			err = rows.Scan(&id, &title, &done)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(id, title, done)
		}
	case "create":
		db, err := sql.Open("sqlite3", dbFile)
		if err != nil {
			log.Fatal("Failed to open database:", err)

		}
		defer db.Close()

		tx, err := db.Begin()

		if err != nil {
			log.Fatal(err)
		}

		query := `INSERT INTO tasks (title, done) VALUES ("GAA", 0)`

		stmt, err := tx.Prepare(query)

		defer stmt.Close()
	default:
		fmt.Println("❌ Command not found.\nUse: todo init")
	}

}
