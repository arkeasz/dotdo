package handlers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func UpdateTask(id int, newTitle string, newDone bool, newDesc string) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal("Failed to open database:", err)

	}
	defer db.Close()

	db.Exec("UPDATE tasks SET title = ?, done = ?, description = ? WHERE id = ?",  newTitle, newDone, newDesc,id)
}
