package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	rootDir := "./test" // Change as needed

	outFile, err := os.Create("all_not_google_files.txt")
	if err != nil {
		fmt.Println("Failed to create output file:", err)
		return
	}
	defer outFile.Close()

	// Walk yyyy folders
	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || !info.IsDir() {
			return nil
		}
		// Check for yyyymm folder (6 digits)
		if len(info.Name()) == 6 && isAllDigits(info.Name()) {
			notGooglePath := filepath.Join(path, "not_google")
			if stat, err := os.Stat(notGooglePath); err == nil && stat.IsDir() {
				filepath.Walk(notGooglePath, func(fpath string, finfo os.FileInfo, ferr error) error {
					if ferr != nil {
						return nil
					}
					if finfo != nil && !finfo.IsDir() {
						fmt.Fprintln(outFile, fpath)
					}
					return nil
				})
			}
			return filepath.SkipDir // Don't descend further into yyyymm
		}
		return nil
	})
}

func isAllDigits(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
