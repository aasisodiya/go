package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Hardcoded directories and output file names

const dirA = "./test/source" // Change as needed
const dirB = "./test/target" // Change as needed

const missingInBFile = "missing_in_dir_B.txt"
const missingInAFile = "missing_in_dir_A.txt"

func main() {
	fmt.Println("--- Directory Comparison Tool ---")
	fmt.Printf("Comparing directory: %s\n", dirA)
	fmt.Printf("With directory: %s\n", dirB)

	// // Create directories for testing if they don't exist
	// if err := createTestDirectories(dirA, dirB); err != nil {
	// 	fmt.Printf("Error setting up test directories: %v\n", err)
	// 	return
	// }

	// 1. Get file bases from both directories
	filesA, err := getFileBases(dirA)
	if err != nil {
		fmt.Printf("Error reading files from %s: %v\n", dirA, err)
		return
	}

	filesB, err := getFileBases(dirB)
	if err != nil {
		fmt.Printf("Error reading files from %s: %v\n", dirB, err)
		return
	}

	// 2. Perform the comparison
	missingB := findMissing(filesA, filesB) // Files in A but not in B
	missingA := findMissing(filesB, filesA) // Files in B but not in A

	// 3. Export results to text files
	if err := exportResults(missingB, missingInBFile, dirA); err != nil {
		fmt.Printf("Error exporting results to %s: %v\n", missingInBFile, err)
	} else {
		fmt.Printf("\n✅ Successfully exported %d missing files (from %s) to %s\n", len(missingB), dirA, missingInBFile)
	}

	if err := exportResults(missingA, missingInAFile, dirB); err != nil {
		fmt.Printf("Error exporting results to %s: %v\n", missingInAFile, err)
	} else {
		fmt.Printf("✅ Successfully exported %d missing files (from %s) to %s\n", len(missingA), dirB, missingInAFile)
	}

	fmt.Println("\n--- Comparison Complete ---")
}

// getFileBases reads the directory and returns a map of file names (without extension)
// to track their presence, ignoring nested directories.
func getFileBases(dirPath string) (map[string]bool, error) {
	bases := make(map[string]bool)

	// Read the directory contents
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		// Ignore nested directories
		if entry.IsDir() {
			continue
		}

		fileName := entry.Name()
		// Get the file name without its extension for comparison
		baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))

		// Convert to lowercase to ensure case-insensitive comparison (optional, but good practice)
		bases[strings.ToLower(baseName)] = true
	}
	return bases, nil
}

// findMissing checks which file bases in source are NOT present in target.
func findMissing(source, target map[string]bool) []string {
	var missing []string
	for base := range source {
		// If the file base is NOT in the target map, it's missing.
		if _, exists := target[base]; !exists {
			missing = append(missing, base)
		}
	}
	return missing
}

// exportResults writes the list of missing file bases to an output file.
func exportResults(missing []string, outputFileName string, sourceDir string) error {
	// Create or truncate the output file
	file, err := os.Create(outputFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	if len(missing) == 0 {
		_, err = file.WriteString(fmt.Sprintf("No files from '%s' were missing in the other directory.\n", sourceDir))
		return err
	}

	// Write the header
	_, err = file.WriteString(fmt.Sprintf("Files present in '%s' but missing in the other directory (comparison by name only):\n", sourceDir))
	if err != nil {
		return err
	}

	// Write each missing file base
	for _, base := range missing {
		_, err = file.WriteString(base + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

// // createTestDirectories is a helper function to set up the environment for testing
// func createTestDirectories(dirs ...string) error {
// 	for _, dir := range dirs {
// 		if err := os.MkdirAll(dir, 0755); err != nil {
// 			return err
// 		}
// 	}

// 	// Create some dummy files for testing the comparison logic
// 	// File 'Image1' is present in both (different extensions) -> NOT MISSING
// 	os.WriteFile(filepath.Join(dirA, "Image1.jpg"), []byte("data"), 0644)
// 	os.WriteFile(filepath.Join(dirB, "Image1.png"), []byte("data"), 0644)

// 	// File 'Video2' is only in A -> MISSING IN B
// 	os.WriteFile(filepath.Join(dirA, "Video2.mp4"), []byte("data"), 0644)

// 	// File 'Audio3' is only in B -> MISSING IN A
// 	os.WriteFile(filepath.Join(dirB, "Audio3.wav"), []byte("data"), 0644)

// 	// Directory inside A (should be ignored)
// 	os.MkdirAll(filepath.Join(dirA, "SubDir"), 0755)
// 	os.WriteFile(filepath.Join(dirA, "SubDir", "ignored.txt"), []byte("data"), 0644)

// 	return nil
// }
