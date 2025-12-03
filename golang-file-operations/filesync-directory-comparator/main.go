// Compare file one by one in two directories and also their subdirectories
// Print the differences

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

	files1 := make(map[string]int64) // key: lowercased relpath, value: size
	files2 := make(map[string]int64)

	collectFiles(dir1, files1)
	collectFiles(dir2, files2)

	f1, _ := os.Create("only_in_dir1.txt")
	defer f1.Close()
	for rel, size1 := range files1 {
		if size2, ok := files2[rel]; !ok {
			fmt.Fprintln(f1, rel)
		} else if size1 != size2 {
			fmt.Fprintf(f1, "%s (DIFFERS SIZE: %d vs %d)\n", rel, size1, size2)
		}
	}

	f2, _ := os.Create("only_in_dir2.txt")
	defer f2.Close()
	for rel, size2 := range files2 {
		if size1, ok := files1[rel]; !ok {
			fmt.Fprintln(f2, rel)
		} else if size1 != size2 {
			fmt.Fprintf(f2, "%s (DIFFERS SIZE: %d vs %d)\n", rel, size2, size1)
		}
	}
}

func collectFiles(root string, files map[string]int64) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() {
			rel, _ := filepath.Rel(root, path)
			rel = strings.ToLower(rel)
			files[rel] = info.Size()
		}
		return nil
	})
}
