package maina

import (
	"fmt"
	"log"
	"os"
	"strings"
	_ "github.com/mattn/go-sqlite3"
	tea "github.com/charmbracelet/bubbletea"
	"todo/pkg/handlers"
)

const dbFile = ".todo/database.db"

type model struct {
	choices  []string
	cursor	 int
	selected map[int]struct{}
	mode     string
	input    textinput.Model
	dbTasks  []handlers.Task
}

func initialModel() model {
	m := model {
		selected: make(map[int]struct{}),
		mode: mode
	}

	if mode == "add" {
		ti := textinput.New()
		ti.Placeholder = "Type the task name"
		ti.Focus()
		m.input = ti
	} else {
		m.dbTasks = handlers.GetAllTasks()

		for  _, task := range m.dbTasks {
			m.choices = append(m.choices, task.Title)
		}
	}

	return
	// tasks := handlers.ShowTasks()

	// tasks_strinng := make([]string, len(tasks))
	// for i, task := range tasks {
	// 	tasks_strinng[i] = task.Title
	// }

	// return model {
	// 	choices: tasks_strinng,
	// 	selected: make(map[int]struct{}),
	// }
}

func (m model) Init() tea.Cmd {
	if m.mode == "add" {
		return textinput.Blink
	}
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices) - 1 {
				m.cursor++
			}
		case "enter", " ":
			if m.mode == "add" {
				task := m.input.Value()

				if task != "" {
					_, err := handlers.AddTask(task)
					if err != nil {
						fmt.Println("Error adding task:", err)
						return m, tea.Quit
					}

					m.dbTasks = handlers.GetAllTasks()
					m.choices = nil
					for _, task := range m.dbTasks {
						m.choices = append(m.choices, task.Title)
					}
				}
				return m, tea.Quit
			}
		}
	}

	if m.mode == "add" {
		var cmd tea.Cmd
		m.input, cmd = m.input.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	switch m.mode {
	case "add":
		return fmt.Sprintf(
			"Añadir nueva tarea:\n\n%s\n\n(Enter para confirmar, Esc para cancelar)",
			m.input.View(),
		)

	case "show":
		s := "📝 Lista de Tareas\n\n"
		for i, task := range m.dbTasks {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			checked := " "
			if task.Done {
				checked = "x"
			}

			s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, task.Title)
		}
		s += "\nPresiona 'q' para salir\n"
		return s

	default:
		return "TODO CLI Application Help:\n\n" +
			"init    - Initialize the database\n" +
			"show    - List all tasks\n" +
			"add     - Add a new task\n" +
			"help    - Show this help message\n"
	}
}

func main()  {


	if len(os.Args) < 2 {
		fmt.Println("❌ You should type a command\nUse: todo init")
		return
	}

	command := strings.ToLower(os.Args[1])

	switch command {
	case "init":
		if err := handlers.InitDB(); err != nil {
			log.Fatal("Error to initialized a database:\n", err)
		}
	case "show":
		p := tea.NewProgram(initialModel())
		if err := p.Start(); err != nil {
			fmt.Println("alas, there's been an error: %v\n", err)
		}
	case "add":
		p := tea.NewProgram(initialModel())
		if err := p.Start(); err != nil {
			fmt.Println("alas, there's been an error: %v\n", err)
		}
	default:
		fmt.Println("❌ Command not found.\nUse: todo help")
	}

}
