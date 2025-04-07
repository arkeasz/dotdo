package handlers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func UpdateTask(id int, newTitle string) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal("Failed to open database:", err)

	}
	defer db.Close()

	result, err := db.Exec("UPDATE tasks SET title = ? WHERE id = ?", newTitle, id)

	if err != nil {
		log.Fatal("Failed to update task:", err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		log.Fatal("Failed to get affected rows:", err)
	}

	if affectedRows == 0 {
		log.Println("No rows were updated. Task may not exist.")
	} else {
		log.Println("Task updated successfully.")
	}
}
