package main

import "github.com/charmbracelet/bubbles/textinput"

type Model struct {
	textInput            textinput.Model
	createNewFileVisible bool
}
