package main

import (
	"fmt"
	"os"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
	"todo/pkg/handlers"
)

const listHeight = 24

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2).Bold(true)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type taskItem struct {
	ID    int
	Label string
	Done  bool
}

func (t taskItem) Title() string       { return t.Label }
func (t taskItem) Description() string {
	d := " "
	if t.Done {
		d = "x"
	}
	return fmt.Sprintf("[%v]", d)
}
func (t taskItem) FilterValue() string { return t.Label }

type model struct {
	Views     []string
	Frame     string
	Choice    taskItem
	TaskList  list.Model
	Quitting  bool
	Selected  bool
	TaskInput textinput.Model
}

func getTaskItems() []list.Item {
	tasks := handlers.GetAllTasks()
	items := make([]list.Item, len(tasks))
	for i, t := range tasks {
		items[i] = taskItem{
			ID:    t.ID,
			Label: t.Title,
			Done:  t.Done,
		}
	}
	return items
}

func initTUI() model {
	ti := textinput.New()
	ti.Placeholder = "Type your task..."
	ti.CharLimit = 156
	ti.Width = 40
	ti.Focus()

	views := []string{"list", "add", "edit", "help", "refresh"}
	frame := views[0]
	items := getTaskItems()

	l := list.New(items, list.NewDefaultDelegate(), 40, listHeight)
	l.Title = "Your Tasks"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return model{
		Views:     views,
		Frame:     frame,
		TaskList:  l,
		TaskInput: ti,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.Quitting = true
			return m, tea.Quit
		case "a":
			if m.Frame == "list" {
				m.Frame = "add"
				m.TaskInput.SetValue("")
				m.TaskInput.Focus()
				return m, textinput.Blink
			}
		case "backspace":
			if m.Frame == "list" {
				if item, ok := m.TaskList.SelectedItem().(taskItem); ok {
					handlers.DeleteTask(item.ID)
					updatedItems := getTaskItems()
					m.TaskList.SetItems([]list.Item{})
					m.TaskList.SetItems(updatedItems)
				}
				return m, nil
			}
		case "tab":
			if m.Frame == "list" {
				if item, ok := m.TaskList.SelectedItem().(taskItem); ok {
					m.Choice = item
					m.TaskInput.SetValue(item.Label)
					m.Selected = true
					m.Frame = "edit"
					m.TaskInput.Focus()
					return m, textinput.Blink
				}
			}
		case "enter":
			if m.Frame == "add" {
				newLabel := m.TaskInput.Value()
				if newLabel != "" {
					handlers.AddTask(newLabel)
					m.TaskList.SetItems(getTaskItems())
					m.Frame = "list"
					m.Selected = false
					m.TaskInput.Blur()
					return m, nil
				}
			}

			if m.Frame == "list" {
				if item, ok := m.TaskList.SelectedItem().(taskItem); ok {
					item.Done = !item.Done
					handlers.UpdateTask(item.ID, item.Label, item.Done)

					items := m.TaskList.Items()
					index := m.TaskList.Index()
					items[index] = item
					m.TaskList.SetItems(items)
					return m, nil
				}
			}

			if m.Frame == "edit" {
				newLabel := m.TaskInput.Value()
				if newLabel != "" {
					handlers.UpdateTask(m.Choice.ID, newLabel, m.Choice.Done)
					m.TaskList.SetItems(getTaskItems())
				}
				m.Frame = "list"
				m.Selected = false
				m.TaskInput.Blur()
				return m, nil
			}
		case "esc":
			if m.Frame == "edit" {
				m.Frame = "list"
				m.Selected = false
				m.TaskInput.Blur()
				return m, nil
			}
		}
	}

	if m.Frame == "list" {
		m.TaskList, cmd = m.TaskList.Update(msg)
	} else {
		m.TaskInput, cmd = m.TaskInput.Update(msg)
	}
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.Quitting {
		return quitTextStyle.Render("Chau 👋")
	}

	switch m.Frame {
	case "list":
		return "\n" + m.TaskList.View()
	case "add":
		return fmt.Sprintf(
			"Add Task:\n\n%s\n\n%s",
			m.TaskInput.View(),
			"(Enter to save, Esc to cancel)",
		)
	case "edit":
		return fmt.Sprintf(
			"Editing Task (ID %d):\n\n%s\n\n%s",
			m.Choice.ID,
			m.TaskInput.View(),
			"(Enter for save, Esc to cancel)",
		)
	default:
		return "\n" + m.TaskList.View()
	}
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "init" {
		handlers.InitDB()
		return
	}

	p := tea.NewProgram(initTUI())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
