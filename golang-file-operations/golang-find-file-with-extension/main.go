package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

var (
	searchPath, fileExtToSearch string
	files                       []string
	opType                      int
)

func main() {
	searchProgram()
}

func searchProgram() {
	fmt.Println("--------Find File With Extension--------")
	//isValidSearchPath := false
	for !IsValidFolderPath(searchPath) {
		fmt.Println("Please Enter the folder path to look into:")
		fmt.Scanln(&searchPath)
	}
	for !IsValidExtension(fileExtToSearch) {
		fmt.Println("Please Enter the extension to look for:")
		fmt.Scanln(&fileExtToSearch)
	}
	fmt.Println("Search Path: ", searchPath, "\nFile Extension To Search:", fileExtToSearch)
	err := filepath.Walk(searchPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if !info.IsDir() && filepath.Ext(path) == fileExtToSearch {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	isValidOperation := false
	for !isValidOperation {
		fmt.Println("Please Select what you want to do with the results:")
		fmt.Println("1. Print the output result")
		fmt.Println("2. Save the result in a file")
		fmt.Scanln(&opType)
		if opType > 0 || opType < 3 {
			isValidOperation = true
		}
	}
	switch opType {
	case 1:
		{
			for _, file := range files {
				fmt.Println(file)
			}
		}
	case 2:
		{
			fmt.Println("Generating file with results")
			if err := WriteResultToFile(files); err == nil {
				fmt.Println("Generated file with results successfully")
			}
		}
	default:
		{
			fmt.Println("Invalid Operation To Perform")
		}
	}
}

// IsValidFolderPath function validates the folderPath and returns true if valid else false (Reference: https://stackoverflow.com/questions/35231846/golang-check-if-string-is-valid-path)
func IsValidFolderPath(folderPath string) (isValid bool) {
	fp := folderPath + string(os.PathSeparator) + "dummy"
	// Check if file already exists
	if _, err := os.Stat(fp); err == nil {
		return true
	}
	// Attempt to create it
	var d []byte
	if err := ioutil.WriteFile(fp, d, 0644); err == nil {
		os.Remove(fp) // And delete it
		return true
	}
	return false
}

// IsValidExtension function validates the extention string and returns true for valid extension else false
func IsValidExtension(ext string) (isValid bool) {
	isValid, _ = regexp.MatchString("\\.[a-zA-Z0-9]+", ext)
	return isValid
}

// WriteResultToFile writes the files list to the output.txt file.
func WriteResultToFile(files []string) error {
	file, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, file := range files {
		fmt.Fprintln(w, file)
	}
	return w.Flush()
}
