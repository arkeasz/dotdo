package ui
import (
	"dotdo/pkg/handlers"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

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
				m.TaskDesc.Blur()
				m.TaskInput.SetValue("")
				m.TaskDesc.SetValue("")
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
					m.TaskDesc.SetValue(item.Desc)
					m.Selected = true
					m.Frame = "edit"
					m.TaskDesc.Blur()
					m.TaskInput.Focus()
					return m, textinput.Blink
				}
			}

			if m.Frame == "edit" {
				m.TaskInput.Blur()
				m.TaskDesc.Focus()
				return m, textarea.Blink
			}

			if m.Frame == "add" {
				m.TaskDesc.Blur()
				m.TaskInput.Blur()
				m.TaskDesc.Focus()
				return m, textarea.Blink
			}
		case "enter":
			if m.Frame == "add" {
				newLabel := m.TaskInput.Value()
				newDesc := m.TaskDesc.Value()
				if len(newDesc) < 1 {
					newDesc = "No description"
				}
				if newLabel != "" {
					typoItem, _ := m.TypoList.SelectedItem().(TypoItem)
					typo := string(typoItem)
					handlers.AddTask(newLabel, newDesc, typo)
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
					handlers.UpdateTask(item.ID, item.Label, item.Done, item.Desc)

					items := m.TaskList.Items()
					index := m.TaskList.Index()
					items[index] = item
					m.TaskList.SetItems(items)
					return m, nil
				}
			}

			if m.Frame == "edit" {
				newLabel := m.TaskInput.Value()
				newDesc := m.TaskDesc.Value()
				if newLabel != "" {
					handlers.UpdateTask(m.Choice.ID, newLabel, m.Choice.Done, newDesc)
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
		cmds = append(cmds, cmd)
	} else {
		m.TaskInput, cmd = m.TaskInput.Update(msg)
		cmds = append(cmds, cmd)
		m.TaskDesc, cmd = m.TaskDesc.Update(msg)
		cmds = append(cmds, cmd)
		m.TypoList, cmd = m.TypoList.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
