package main

import (
	"fmt"
	"log"
	"os"
)

// ListFilesOnly reads the directory at the given path and prints the names of
// only the regular files within it.
func ListFilesOnly(dirPath string) {
	// Read the directory contents
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Error reading directory %s: %v", dirPath, err)
	}

	fmt.Printf("Files in '%s':\n", dirPath)

	foundFiles := false
	// Iterate through the entries
	for _, entry := range entries {
		// Check if the entry is NOT a directory
		if !entry.IsDir() {
			// Get the name of the file
			fmt.Println(entry.Name())
			foundFiles = true
		}
	}

	if !foundFiles {
		fmt.Println("(No regular files found)")
	}
}

func main() {
	// --- IMPORTANT ---
	// Replace "path/to/your/folder" with the actual path
	// e.g., "." for the current directory, or "/home/user/documents"
	folderPath := "./"

	ListFilesOnly(folderPath)
}
