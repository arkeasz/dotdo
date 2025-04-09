package ui

import (
	"fmt"
	"github.com/olekukonko/ts"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
	"dotdo/pkg/handlers"
)


const listHeight = 24

var (
	titleStyle        = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63"))
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4).Foreground(lipgloss.Color("240"))
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("204")).Bold(true)
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

func getSize() (int, int) {
	size, err := ts.GetSize()
	if err != nil {
		panic(err)
	}
	rows := size.Row()
	cols := size.Col()
	return rows, cols
}

var (
	totalWidth, totalHeight = getSize()
	listStyle = lipgloss.NewStyle().
        Width((totalHeight-10) / 2).
        BorderStyle(lipgloss.NormalBorder()).
        BorderForeground(lipgloss.Color("62")).
        Padding(1, 2)

    sidePanelStyle = lipgloss.NewStyle().
        Width((totalHeight-10) / 2).
        BorderStyle(lipgloss.NormalBorder()).
        BorderForeground(lipgloss.Color("63")).
        Padding(1, 2)
)

type taskItem struct {
	ID    int
	Label string
	Done  bool
	Typo  string
	Desc  string
}

func (t taskItem) Title() string       { return t.Label }
func (t taskItem) Description() string {
	d := " "
	if t.Done {
		d = "x"
	}
	return fmt.Sprintf("[%v] (%v)", d, t.Typo)
}
func (t taskItem) FilterValue() string { return t.Label }

type model struct {
	Views     []string
	Frame     string
	Choice    taskItem
	TaskList  list.Model
	TypoList  list.Model
	Quitting  bool
	Selected  bool
	TaskInput textinput.Model
	TaskDesc  textarea.Model
}

type TypoItem string
func (t TypoItem) Title() string       { return string(t) }
func (t TypoItem) Description() string { return "" }
func (t TypoItem) FilterValue() string { return string(t) }
func getTaskItems() []list.Item {
	tasks := handlers.GetAllTasks()
	items := make([]list.Item, len(tasks))
	for i, t := range tasks {
		items[i] = taskItem{
			ID:    t.ID,    // ID
			Label: t.Title, // Title
			Done:  t.Done,  // Done
			Typo:  t.Typo,  // Typo
			Desc:  t.Desc,  // Description
		}
	}
	return items
}

func Ran() model {
	ti := textinput.New()
	ti.Placeholder = "Type your task..."
	ti.CharLimit = 130
	ti.Width = 40
	ti.Focus()
	ta := textarea.New()
	ta.Placeholder = "Type your task description..."
	ta.ShowLineNumbers = false


	views := []string{"list", "add", "edit", "help"}
	frame := views[0]
	items := getTaskItems()

	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle = selectedItemStyle
	delegate.Styles.SelectedDesc = selectedItemStyle


	l := list.New(items, delegate, totalHeight - 20, listHeight)
	l.Title = "Your Tasks"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	typoItems := []list.Item{
		TypoItem("todo"),
		TypoItem("bug"),
		TypoItem("feature"),
		TypoItem("docs"),
		TypoItem("refactor"),
	}

	typoDelegate := list.NewDefaultDelegate()
	typoDelegate.Styles.SelectedTitle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("212")).Bold(true)

	t := list.New(typoItems, typoDelegate, totalHeight - 10, 10)
	t.Title = "Select a Typo"
	t.SetShowHelp(false)
	t.SetShowPagination(false)
	t.SetShowStatusBar(false)
	return model{
		Views:     views,
		Frame:     frame,
		TaskList:  l,
		TaskInput: ti,
		TaskDesc:  ta,
		TypoList:  t,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
