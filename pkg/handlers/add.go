package handlers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func AddTask(task string, desc string) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal("Failed to open database:", err)

	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO tasks (title, done, description) VALUES (?, 0, ?)", task, desc)

	if err != nil {
		log.Fatal("Failed to insert task:", err)
	}
}
