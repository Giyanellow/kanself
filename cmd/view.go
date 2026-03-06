package main

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m Model) View() tea.View {
	if !m.loaded {
		return tea.NewView("Loading...")
	}

	// Define styles
	focusedStyle := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("62"))
	normalStyle := lipgloss.NewStyle().Border(lipgloss.HiddenBorder())

	// Create a slice to hold the styled column views
	columns := []string{}

	// Loop through each list
	for i := range m.list {
		var styledView string
		if status(i) == m.focused {
			styledView = focusedStyle.Render(m.list[i].View())
		} else {
			styledView = normalStyle.Render(m.list[i].View())
		}
		columns = append(columns, styledView)
	}

	// Join them horizontally
	allColumns := lipgloss.JoinHorizontal(lipgloss.Left, columns...)
	help := m.help.View(m.keys)
	allColumnsWithHelp := lipgloss.JoinVertical(lipgloss.Left, allColumns, help)
	return tea.NewView(allColumnsWithHelp)
}
