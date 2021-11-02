package main

import (
	"fmt"
	"os"
)

var (
	filePath string
)

func main() {
	FileMetaDataProgram()
}

// FileMetaDataProgram will help you to get and print the metadata for given filepath
func FileMetaDataProgram() {
	fmt.Println("----------File Meta Data Program----------")
	fmt.Println("Enter the file path for which to get metadata:")
	fmt.Scanln(&filePath)
	fmt.Println("Getting Metedata for File:", filePath)
	//The file has to be opened first
	f, _ := os.Open(filePath)
	defer f.Close()
	// The file descriptor (File*) has to be used to get metadata
	fi, err := f.Stat()
	// The file can be closed
	if err != nil {
		fmt.Println(err)
		return
	}
	// fi is a fileInfo interface returned by Stat
	fmt.Println("File Name:", fi.Name())
	fmt.Println("File Size (bytes):", fi.Size())
	fmt.Println("Is Directory?:", fi.IsDir())
	fmt.Println("Modification Time:", fi.ModTime())
}
