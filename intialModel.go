package main

import (
	"os"

	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) Init() tea.Cmd {
	directoryLocation = getDirectoryLocation()
	if directoryLocation == "" {
		fmt.Println("Please set the NOTES_DIRECTORY environment variable in a .env file.")
		return tea.Quit
	}
	if _, err := os.Stat(directoryLocation); os.IsNotExist(err) {
		fmt.Printf("The directory %s does not exist. Please create it or set a valid NOTES_DIRECTORY.\n", directoryLocation)
		return tea.Quit
	}
	return nil
}
func intialSetupTextInput() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "What you would like to call your note?"
	ti.CharLimit = 200
	ti.Width = 50
	ti.Focus()
	ti.Cursor.Style = cursorStyle
	ti.PromptStyle = cursorStyle
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	return ti
}
func intialTextArea() textarea.Model {
	ti := textarea.New()
	ti.Placeholder = "Enter ur note here..."
	ti.Focus()
	return ti
}

func initialModel() Model {
	ti := intialSetupTextInput()
	ta := intialTextArea()
	return Model{
		textInput:            ti,
		createNewFileVisible: false,
		textarea:             ta,
	}
}
