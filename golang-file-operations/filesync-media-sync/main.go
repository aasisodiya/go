package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Starting Execution...")
	backupDir := "./test/source"
	externalDir := "./test/target"

	backupFiles := make(map[string][]string)   // key: lowercased filename, value: list of relpaths
	externalFiles := make(map[string][]string) // key: lowercased filename, value: list of relpaths

	// Collect media files from yyyymm and yyyymm/not_google in backup
	collectMonthAndNotGoogleFiles(backupDir, backupFiles)
	// Collect media files from yyyymm and yyyymm/not_google in external drive
	collectMonthAndNotGoogleFiles(externalDir, externalFiles)

	// Write missing in external drive to file
	f1, err := os.Create("missing_in_external_drive.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating missing_in_external_drive.txt: %v\n", err)
		return
	}
	defer f1.Close()
	for name, relpaths := range backupFiles {
		if _, ok := externalFiles[name]; !ok {
			for _, rel := range relpaths {
				fmt.Fprintln(f1, rel)
			}
		}
	}

	// Write missing in backup to file
	f2, err := os.Create("missing_in_backup.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating missing_in_backup.txt: %v\n", err)
		return
	}
	defer f2.Close()
	for name, relpaths := range externalFiles {
		if _, ok := backupFiles[name]; !ok {
			for _, rel := range relpaths {
				fmt.Fprintln(f2, rel)
			}
		}
	}
}

// Collect files from all yyyymm and yyyymm/not_google folders, all formats, no skipping
func collectMonthAndNotGoogleFiles(root string, files map[string][]string) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
if err != nil {
		fmt.Fprintf(os.Stderr, "error walking path %q: %v\n", path, err)
		return nil
	}
		}
		if !info.IsDir() {
			return nil
		}
		// Check if this is a yyyymm folder (6 digits)
		if len(info.Name()) == 6 && isAllDigits(info.Name()) {
			// Collect files in yyyymm
			filepath.Walk(path, func(fpath string, finfo os.FileInfo, ferr error) error {
if ferr != nil {
				fmt.Fprintf(os.Stderr, "error walking path %q: %v\n", fpath, ferr)
				return nil
			}
				}
				// Only collect files directly under yyyymm (not in subfolders)
				if finfo != nil && !finfo.IsDir() && filepath.Dir(fpath) == path {
					key := strings.ToLower(finfo.Name())
					rel, _ := filepath.Rel(root, fpath)
					files[key] = append(files[key], rel)
				}
				return nil
			})
			// Collect files in yyyymm/not_google if exists
			notGooglePath := filepath.Join(path, "not_google")
			if stat, err := os.Stat(notGooglePath); err == nil && stat.IsDir() {
				filepath.Walk(notGooglePath, func(fpath string, finfo os.FileInfo, ferr error) error {
if ferr != nil {
					fmt.Fprintf(os.Stderr, "error walking path %q: %v\n", fpath, ferr)
					return nil
				}
					}
					if finfo != nil && !finfo.IsDir() {
						key := strings.ToLower(finfo.Name())
						rel, err := filepath.Rel(root, fpath)
if err != nil {
					fmt.Fprintf(os.Stderr, "error getting relative path for %q from root %q: %v\n", fpath, root, err)
					return nil
				}
						files[key] = append(files[key], rel)
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

// func isMediaFile(name string) bool {
// 	ext := strings.ToLower(filepath.Ext(name))
// 	switch ext {
// 	case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".mp4", ".mov", ".avi", ".mkv", ".heic", ".tiff", ".webp":
// 		return true
// 	default:
// 		fmt.Println("Skipping file with extension:", ext, "File-name:", name)
// 		return false
// 	}
// }
