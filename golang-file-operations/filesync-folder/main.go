package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// The mediaExtensions map defines the standard extensions and their equivalents.
var mediaExtensions = map[string]string{
	".heic": ".jpg",
	".jpeg": ".jpg",
	".jpg":  ".jpg",
	".png":  ".png",
	".gif":  ".gif",
	".bmp":  ".bmp",
	".tiff": ".tiff",
	".webp": ".webp",

	".mov": ".mp4",
	".mp4": ".mp4",
	".avi": ".avi",
	".mkv": ".mkv",
}

// normalizeFilename takes a filename and returns a normalized version.
// This handles case-insensitivity and standardizes file extensions
// based on the mediaExtensions map. For example, "image.HEIC" becomes "image.jpg".
func normalizeFilename(filename string) (string, bool) {
	fmt.Println(filename)
	// Get the file's base name and extension in lowercase.
	ext := strings.ToLower(filepath.Ext(filename))
	base := strings.TrimSuffix(filename, filepath.Ext(filename))

	// Check if the extension is a recognized media type.
	normalizedExt, ok := mediaExtensions[ext]
	if !ok {
		// If the extension is not a recognized media type, return the original filename
		// and indicate that it was not a media file.
		return filename, false
	}

	// Construct the normalized filename.
	normalizedName := strings.ToLower(base) + normalizedExt
	return normalizedName, true
}

// getFilesMap walks through a directory and creates a map of normalized filenames
// to their original filenames.
func getFilesMap(dir string) (map[string]string, error) {
	filesMap := make(map[string]string)

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	fmt.Println(len(entries))

	for _, entry := range entries {
		if !entry.IsDir() {
			normalizedName, isMedia := normalizeFilename(entry.Name())
			if isMedia {
				filesMap[normalizedName] = entry.Name()
			}
		}
	}

	return filesMap, nil
}

// findMissingFiles compares two file maps and returns a slice of original filenames
// present in mapA but not mapB.
func findMissingFiles(mapA, mapB map[string]string) []string {
	var missingFiles []string
	for normalizedName, originalName := range mapA {
		// Check for the presence of the normalized name in the other map.
		if _, ok := mapB[normalizedName]; !ok {
			missingFiles = append(missingFiles, fmt.Sprintf("move \"%s\" \".\\not_google\\%s\"", originalName, originalName))
		}
	}
	return missingFiles
}

// writeToFile writes a slice of strings to a specified file, with each string on a new line.
func writeToFile(filename string, files []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filename, err)
	}
	defer file.Close()

	for _, f := range files {
		_, err := fmt.Fprintln(file, f)
		if err != nil {
			return fmt.Errorf("failed to write to file %s: %w", filename, err)
		}
	}

	return nil
}

func main() {
	// Check for the correct number of command-line arguments.
	// if len(os.Args) != 3 {
	// 	fmt.Println("Usage: go run compare_media_files.go <folder1> <folder2>")
	// 	os.Exit(1)
	// }

	folder1 := "./test/source/202501"
	folder2 := "./test/target/202501"

	log.Printf("Scanning files in folder 1: %s", folder1)
	filesInFolder1, err := getFilesMap(folder1)
	if err != nil {
		log.Fatalf("Error reading folder 1: %v", err)
	}

	log.Printf("Scanning files in folder 2: %s", folder2)
	filesInFolder2, err := getFilesMap(folder2)
	if err != nil {
		log.Fatalf("Error reading folder 2: %v", err)
	}

	fmt.Println(len(filesInFolder1), len(filesInFolder2))

	// Find files missing from each folder.
	missingFromFolder2 := findMissingFiles(filesInFolder1, filesInFolder2)
	missingFromFolder1 := findMissingFiles(filesInFolder2, filesInFolder1)

	// Write the lists of missing files to separate files.
	err = writeToFile("missing_from_folder1.txt", missingFromFolder1)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
	err = writeToFile("missing_from_folder2.txt", missingFromFolder2)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}

	fmt.Println("\nComparison complete.")
	fmt.Println("Lists of missing files have been saved to the following files:")
	fmt.Println("- missing_from_folder1.txt")
	fmt.Println("- missing_from_folder2.txt")
}
