package handlers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func DeleteTask(id int) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	db.Exec("DELETE FROM tasks WHERE id = ?", id)

}
