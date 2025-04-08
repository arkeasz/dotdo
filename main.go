package main

import (
	"fmt"
	"os"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	tea "github.com/charmbracelet/bubbletea"
	"dotdo/pkg/handlers"
	"dotdo/pkg/ui"
)

func main() {
	if len(os.Args) > 1 {
		command := os.Args[1]

		switch command {
		case "init":
			handlers.InitDB()
			return
		case "drop":
			if err := handlers.DeleteDBFile(); err != nil {
				log.Fatal(err)
			}
			if err := handlers.DropTable(); err != nil {
				log.Fatal(err)
			}
			return
		}
	}

	db, err := sql.Open("sqlite3", ".todo/database.db")
	if err != nil {
		fmt.Println("Type dotdo help")
		return
	}
	defer db.Close()

	p := tea.NewProgram(ui.Ran())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
