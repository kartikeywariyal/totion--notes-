# Totion - Terminal Notes Application

A lightweight, terminal-based note-taking application built with Go. Totion provides a simple and efficient way to create, edit, and manage notes directly from your terminal.

## Features

- 📝 Create and edit notes in the terminal
- 🎨 Beautiful TUI (Text User Interface) powered by Bubble Tea
- ⚡ Fast and lightweight
- 🔧 Built with Go for excellent performance

## Tech Stack

- **Go** 1.25.5
- **Bubble Tea** - Terminal UI framework
- **Lipgloss** - Terminal styling and formatting
- **Bubbles** - Reusable Bubble Tea components

## Project Structure

- `main.go` - Application entry point
- `model.go` - Core data model for the application
- `views.go` - UI views and rendering
- `update.go` - Update logic and event handling
- `creation.go` - File creation functionality
- `intialModel.go` - Model initialization
- `go.mod` - Go module dependencies

## Installation

```bash
git clone https://github.com/kartikeywariyal/totion.git
cd totion
go install
```

## Usage

Run the application:

```bash
go run main.go
```

Or if installed:

```bash
totion
```

## Dependencies

- `github.com/charmbracelet/bubbletea` - Terminal UI framework
- `github.com/charmbracelet/lipgloss` - Terminal styling
- `github.com/charmbracelet/bubbles` - Reusable UI components

## License

MIT License

## Author

Kartikey Wariyal
