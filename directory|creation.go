package main

import (
	"fmt"
	"log"
	"os"

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
