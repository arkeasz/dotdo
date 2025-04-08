package ui
import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	if m.Quitting {
		return quitTextStyle.Render("Chau 👋")
	}

	listPanel := listStyle.Render(m.TaskList.View())
	var sidePanel string

	switch m.Frame {
    case "add":
		sidePanel = sidePanelStyle.Render(
			fmt.Sprintf("Add Task:\n\n%s\n\n%s\n\n%s\n\n%s",
				m.TaskInput.View(),
				m.TaskDesc.View(),
				m.TypoList.View(),
				"(Enter to save, Esc to cancel)"),
		)
    case "edit":
        sidePanel = sidePanelStyle.Render(
            fmt.Sprintf("Editing Task (ID %d):\n\n%s\n\n%s\n\n%s",
                m.Choice.ID,
                m.TaskInput.View(),
				m.TaskDesc.View(),
                "(Enter for save, Esc to cancel)"),
        )
    default:
		if item, ok := m.TaskList.SelectedItem().(taskItem); ok {
			sidePanel = sidePanelStyle.Render(
				fmt.Sprintf("%s\n\nDetails:\n%s",
					item.Label,
					item.Desc,
				),
			)
		} else {
			sidePanel = sidePanelStyle.Render("No task selected")
		}
    }

    return "\n" + lipgloss.JoinHorizontal(
        lipgloss.Top,
        listPanel,
        sidePanel,
    ) + "\n"
}
