package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	_ "github.com/mattn/go-sqlite3"
	tea "github.com/charmbracelet/bubbletea"
	"todo/pkg/handlers"
)

type model struct {}

func initialModel(mode string) model {

}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case: "ctrl+c"
		}
	}
	return m, nil
}

func (m model) View() string {
	return "Hello, world"
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("err: %w", err)
	}

	defer f.Close()

	p := tea.NewProgram(model )

	// if len(os.Args) < 2 {
	// 	fmt.Println("❌ You should type a command\nUse: todo init")
	// 	return
	// }

	// command := strings.ToLower(os.Args[1])

	// switch command {
	// case "init":
	// 	if err := handlers.InitDB(); err != nil {
	// 		log.Fatal("Error to initialized a database:\n", err)
	// 	}
	// case "show":
	// 	p := tea.NewProgram(initialModel("show"))
	// 	if err := p.Run(); err != nil {
	// 		fmt.Println("GAAAAAAAAAAAAAAA")
	// 	}
	// }
}
