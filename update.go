package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd = nil
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+n":
			m.createNewFileVisible = true
			return m, nil
		case "enter":
			if m.createNewFileVisible {
				filename := m.textInput.Value()
				if filename != "" {
					filepath := fmt.Sprintf("%s/%s.txt", directoryLocation, filename)
					err := createNewNoteFile(filepath, &m)
					if err != nil {
						fmt.Println("Error creating file:", err)
					}
					m.createNewFileVisible = false
					m.textInput.SetValue("")
				}
			}
		case "ctrl+l":
			// Future implementation for listing notes
		case "ctrl+s":
			m.fileDescriptor.Write([]byte(m.textarea.Value()))
		case "esc":
			m.createNewFileVisible = false
			m.textInput.SetValue("")
			m.textarea.SetValue("")
			m.fileDescriptor = nil
		}
	}
	if m.createNewFileVisible {
		m.textInput, cmd = m.textInput.Update(msg)
	}
	if m.fileDescriptor != nil {
		m.textarea, cmd = m.textarea.Update(msg)
	}
	return m, cmd
}
