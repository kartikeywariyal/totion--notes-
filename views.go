package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	welcomePage := "Welcome to Kartikey Totion 🧠"
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("16")).
		Background(lipgloss.Color("205")).Align(lipgloss.Center).
		Padding(0, 2)

	welcomePage = style.Render(welcomePage)
	help := "Ctrl+N  New  ·  Ctrl+L  List  ·  Ctrl+S  Save  ·  Esc  Back  ·  Ctrl+C/ Quit"
	view := ""
	if m.createNewFileVisible {
		view = m.textInput.View()
	}
	return fmt.Sprintf("\n%s\n\n%s\n\n%s\n", welcomePage, view, help)
}
