// I want a golang code that will print the number of files in a folder and its subfolders and also print the total number of all files in those folders inluding subfolders.
package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

// fileCounts stores the file count for a specific directory path.
type fileCounts map[string]int

// walkDirectory recursively walks the directory and populates the fileCounts map.
// It returns the total number of files found (including subfolders) and any error.
func walkDirectory(root string) (int, fileCounts, error) {
	counts := make(fileCounts)
	totalFiles := 0

	// filepath.WalkDir is a robust function for recursive directory traversal.
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		// If there was an error accessing the path, print it and continue (or return err to stop).
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return nil // Continue the walk to other paths
		}

		// Check if the current entry is a directory.
		if d.IsDir() {
			// Initialize the count for this directory path.
			// This ensures that even directories with 0 files are in the map.
			counts[path] = 0
			return nil // Continue traversal into this directory
		}

		// If it's a regular file or a "desktop.ini" file: 
		if d.Type().IsRegular() && filepath.Base(path) != "desktop.ini" {
			totalFiles++

			// Get the directory containing the file.
			dir := filepath.Dir(path)

			// Increment the file count for its immediate parent directory.
			counts[dir]++
		}

		return nil
	})

	// Add the count for the root directory itself, which might not be in the map
	// if it contains no files but only subdirectories (initialized with 'counts[path] = 0').
	// It's safer to ensure the root path is checked again.
	if _, ok := counts[root]; !ok {
		// If the root wasn't processed as a directory in the walk (e.g., if it's the
		// starting directory that WalkDir immediately starts processing content for),
		// we still initialize its count to 0 if it's not present.
		// However, filepath.WalkDir starts with the root itself, so it should be there.
		// We'll rely on the logic inside the walk function to handle this for the directories.
	}

	return totalFiles, counts, err
}

func main() {
	// // 1. Get the target directory from command-line arguments.
	// if len(os.Args) < 2 {
	// 	log.Fatal("Usage: go run your_script_name.go <directory_path>")
	// }
	targetDir := "./test/" // Hardcoded path

	// 2. Resolve the path to an absolute path for clear output.
	absDir, err := filepath.Abs(targetDir)
	if err != nil {
		log.Fatalf("Error resolving absolute path: %v", err)
	}

	// 3. Perform the recursive file count.
	total, counts, err := walkDirectory(absDir)
	if err != nil {
		log.Fatalf("Error walking directory: %v", err)
	}

	// 4. Print the file count for each folder and subfolder.
	fmt.Printf("--- File Count Per Folder (in %s) ---\n", absDir)

	// Create a slice of paths to sort for deterministic output.
	var dirs []string
	for dir := range counts {
		dirs = append(dirs, dir)
	}

	// Sort the directory paths alphabetically.
	// This helps with readability, especially for nested folders.
	for i := 0; i < len(dirs); i++ {
		for j := i + 1; j < len(dirs); j++ {
			if dirs[i] > dirs[j] {
				dirs[i], dirs[j] = dirs[j], dirs[i]
			}
		}
	}

	for _, dir := range dirs {
		// The path printed is relative to the root for better output readability,
		// but the map key is the absolute path.
		relPath, _ := filepath.Rel(absDir, dir)
		if relPath == "." {
			// Special case for the root directory itself
			fmt.Printf("Root Directory (%s): %d files\n", absDir, counts[dir])
		} else {
			fmt.Printf("Subfolder: %s: %d files\n", relPath, counts[dir])
		}
	}

	// 5. Print the total file count.
	fmt.Println("-----------------------------------------------------")
	fmt.Printf("Total files in all folders and subfolders: %d\n", total)
}
