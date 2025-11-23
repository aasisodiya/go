// So i have a backup of my media files specifically photos and videos
// and i want to compare it with the files that are in my external drive
// and find out which files are missing in either of the folders
// but you need to also check if those files are present in nested directories
// and if they are not present in nested directories then list them out

// Now my folder structure is like this, yyyy is folder and inside that i have a folder named yyyymm for all the months of that year and inside that i have another folder named not_google
// now while comparing the files in nested directories i am referring to not_google folder
// so basically you want to loop the code on folder yyyy and subfolder yyyymm and then compare the files in yyyymm and also yyyymm/not_google

// What i need is to compare the files in yyyymm and yyyymm/not_google of backup with the files in external drive
// and find out which files are missing in either of the folders
// also while comparing the files consider all file formats don't skip any file formats
// also let say a file is present in yyyymm of backup and in not_google of external drive then consider it as present in both the folders
// and vice versa
// and also i want to hardcode the paths in the code itself don't take input from command line arguments

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
			fmt.Println(err)
			return nil
		}
		if !info.IsDir() {
			return nil
		}
		// Check if this is a yyyymm folder (6 digits)
		if len(info.Name()) == 6 && isAllDigits(info.Name()) {
			// Collect files in yyyymm
			filepath.Walk(path, func(fpath string, finfo os.FileInfo, ferr error) error {
				if ferr != nil {
					fmt.Println(ferr)
					return nil
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
						fmt.Println(ferr)
						return nil
					}
					if finfo != nil && !finfo.IsDir() {
						key := strings.ToLower(finfo.Name())
						rel, err := filepath.Rel(root, fpath)
						if err != nil {
							fmt.Println(err)
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
// 		fmt.Println("Skipping file with extension:", ext, "Filename:", name)
// 		return false
// 	}
// }
