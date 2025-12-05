package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir1 := "./test/source" // Change as needed
	dir2 := "./test/target" // Change as needed

	folders1 := make(map[string]string) // key: lowercased relpath, value: full path
	folders2 := make(map[string]string)

	collectFolders(dir1, folders1)
	collectFolders(dir2, folders2)

	f1, _ := os.Create("only_in_dir1_folders.txt")
	defer f1.Close()
	for rel, full := range folders1 {
		if _, ok := folders2[rel]; !ok {
			fmt.Fprintln(f1, full)
		}
	}

	f2, _ := os.Create("only_in_dir2_folders.txt")
	defer f2.Close()
	for rel, full := range folders2 {
		if _, ok := folders1[rel]; !ok {
			fmt.Fprintln(f2, full)
		}
	}
}

func collectFolders(root string, folders map[string]string) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			rel, _ := filepath.Rel(root, path)
			rel = strings.ToLower(rel)
			folders[rel] = path // store full path
			fmt.Println("Found folder:", path)
		}
		return nil
	})
}
