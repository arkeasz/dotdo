package handlers

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func GetAllTasks() []Task {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	tasks := []Task{}

	query := `SELECT id, title, done FROM tasks;`

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var task Task

		err = rows.Scan(&task.ID, &task.Title, &task.Done)

		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, task)
	}

	return tasks
}
