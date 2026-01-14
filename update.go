package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd = nil
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+n":
			m.createNewFileVisible = true
			return m, nil
		case "enter":
			if m.createNewFileVisible {
				// create a new file with the name from text input
			}
		}
	}
	if m.createNewFileVisible {
		m.textInput, cmd = m.textInput.Update(msg)
	}
	return m, cmd
}
