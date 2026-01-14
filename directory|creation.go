package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/joho/godotenv"
)

var directoryLocation string

func getDirectoryLocation() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	directoryLocation := os.Getenv("NOTES_DIRECTORY")
	return directoryLocation
}
func createNewNoteFile(filePath string, m *Model) error {
	if _, err := os.Stat(filePath); err == nil {
		return fmt.Errorf("file %s already exists", filePath)
	}
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	m.fileDescriptor = file
	return nil
}

// get List for Notes
func getListOfNotes() ([]list.Item, error) {
	files, err := os.ReadDir(directoryLocation)
	if err != nil {
		return nil, fmt.Errorf("error while reading directory %s: %w", directoryLocation, err)
	}
	var items []list.Item
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		Info, _ := file.Info()
		items = append(items, item{
			title: file.Name(),
			desc:  Info.ModTime().Format("2006:01:02 15:04:05"),
		})
	}

	return items, nil
}
