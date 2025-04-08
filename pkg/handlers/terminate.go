package handlers

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/mattn/go-sqlite3"
)

func DropTable() error {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	query := `DROP TABLE IF EXISTS tasks;`
	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to drop table: %v", err)
	}

	fmt.Println("✅ Table dropped successfully")
	return nil
}

func DeleteDBFile() error {
	fmt.Printf("Deleting database file: %s\n", dbFile)
	err := os.Remove(dbFile)
	if err != nil {
		return fmt.Errorf("failed to delete database file: %v", err)
	}

	todoDir := ".todo"
	err = os.Remove(todoDir)
	if err != nil {
		fmt.Printf("Note: Could not remove directory %s (may not be empty or already deleted)\n", todoDir)
	}

	fmt.Println("✅ Database file and directory removed successfully")
	return nil
}
