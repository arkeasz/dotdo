package handlers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func AddTask(task string) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal("Failed to open database:", err)

	}
	defer db.Close()

	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	query := `INSERT INTO tasks (title, done) VALUES (?, 0)`

	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}


}
