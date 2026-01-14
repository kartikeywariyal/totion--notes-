package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd = nil
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+n":
			m.createNewFileVisible = true
			return m, nil
		case "ctrl+d":
			if m.showList {
				selectedItem, ok := m.list.SelectedItem().(item)
				if ok {
					filepath := fmt.Sprintf("%s/%s", directoryLocation, selectedItem.title)
					err := os.Remove(filepath)
					if err != nil {
						log.Fatal("Error deleting file", err)
					}
					item, _ := getListOfNotes()
					finalList := list.New(item, list.NewDefaultDelegate(), 0, 0)
					m.list.SetItems(finalList.Items())
					m.showList = true
				}
				return m, nil
			}
		case "enter":
			if m.createNewFileVisible {
				filename := m.textInput.Value()
				if filename != "" {
					filepath := fmt.Sprintf("%s/%s.txt", directoryLocation, filename)
					err := createNewNoteFile(filepath, &m)
					if err != nil {
						m.textInput.SetValue("")
						fmt.Println("Error creating file:", err)
					} else {
						m.createNewFileVisible = false
						m.textInput.SetValue("")
					}

				}
			}
			if m.showList {
				selectedItem, ok := m.list.SelectedItem().(item)
				if ok {
					filepath := fmt.Sprintf("%s/%s", directoryLocation, selectedItem.title)
					fd, err := os.OpenFile(filepath, os.O_RDWR, 0644)
					if err != nil {
						log.Fatal("Error opening file", err)
					} else {
						m.fileDescriptor = fd
					}
					m.showList = false
					data, err := os.ReadFile(filepath)
					if err != nil {
						log.Fatal("Error reading file", err)
					}
					m.textarea.SetValue(string(data))
					m.textarea.Focus()
				}
				return m, nil

			}
		case "ctrl+l":
			m.list = initialModel().list
			m.list.SetSize(80, 20)
			m.showList = true
		case "ctrl+s":
			m.fileDescriptor.Write([]byte(m.textarea.Value()))
		case "esc":
			m.createNewFileVisible = false
			m.textInput.SetValue("")
			m.textarea.SetValue("")
			m.fileDescriptor = nil
			m.showList = false
		}
	}
	if m.createNewFileVisible {
		m.textInput, cmd = m.textInput.Update(msg)
	}
	if m.fileDescriptor != nil {
		m.textarea, cmd = m.textarea.Update(msg)
	}
	if m.showList {
		m.list, cmd = m.list.Update(msg)
	}
	return m, cmd
}
