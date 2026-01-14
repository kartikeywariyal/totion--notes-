package main

import (
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
)

type item struct {
	title, desc string
}
type Model struct {
	textInput            textinput.Model
	createNewFileVisible bool
	fileDescriptor       *os.File
	textarea             textarea.Model
	list                 list.Model
	showList             bool
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }
