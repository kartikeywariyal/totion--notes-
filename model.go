package main

import (
	"os"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
)

type Model struct {
	textInput            textinput.Model
	createNewFileVisible bool
	fileDescriptor       *os.File
	textarea             textarea.Model
}
