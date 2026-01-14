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
	help := "Ctrl+N  New  ·  Ctrl+L  List  ·  Ctrl+S  Save ·  Ctrl+D   Delete  ·  Esc  Back  ·  Ctrl+C  Quit"
	view := ""
	// newFile view
	if m.createNewFileVisible {
		view = m.textInput.View()
	}

	// Writing Notes view

	if m.fileDescriptor != nil {
		view = m.textarea.View()
	}
	// List View

	if m.showList {
		view = m.list.View()
	}

	return fmt.Sprintf("\n%s\n\n\n%s\n\n\n%s\n", welcomePage, view, help)
}
