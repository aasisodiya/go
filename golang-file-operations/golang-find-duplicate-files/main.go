package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	rootDir := "./test/target" // Change as needed

	outFile, err := os.Create("duplicates_in_not_google.txt")
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
			yyyymmPath := path
			notGooglePath := filepath.Join(yyyymmPath, "not_google")
			// Only proceed if not_google exists
			if stat, err := os.Stat(notGooglePath); err == nil && stat.IsDir() {
				// Collect files in yyyymm (not in subfolders)
				filesInMonth := make(map[string]string)
				filepath.Walk(yyyymmPath, func(fpath string, finfo os.FileInfo, ferr error) error {
					if ferr != nil {
						return nil
					}
					if finfo != nil && !finfo.IsDir() && filepath.Dir(fpath) == yyyymmPath {
						filesInMonth[strings.ToLower(finfo.Name())] = fpath
					}
					return nil
				})
				// Collect files in not_google
				filesInNotGoogle := make(map[string]struct{})
				filepath.Walk(notGooglePath, func(fpath string, finfo os.FileInfo, ferr error) error {
					if ferr != nil {
						return nil
					}
					if finfo != nil && !finfo.IsDir() {
						filesInNotGoogle[strings.ToLower(finfo.Name())] = struct{}{}
					}
					return nil
				})
				// Find duplicates
				for fname, fullpath := range filesInMonth {
					if _, ok := filesInNotGoogle[fname]; ok {
						fmt.Println("Duplicate file:", fname, "Full path:", fullpath)
						fmt.Fprintln(outFile, fullpath)
					}
				}
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
