package archieve

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

// ZipFiles compresses one or many files into a single zip archive file.
// Param 1: filename is the output zip file's name.
// Param 2: files is a list of files to add to the zip.
func ZipFiles(filename string, files []string) error {

	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {

		fileToZip, err := os.Open(file)
		if err != nil {
			fmt.Println("File not found", err) // Remove the file.
			newZipFile.Close()
			err1 := os.Remove(filename)
			if err1 != nil {
				return err1
			}
			fmt.Println("Removed file:", filename)
			return err
		}
		if err = AddFileToZip(zipWriter, file, fileToZip); err != nil {
			return err
		}
	}
	return nil
}

// AddFileToZip will add the file to zip archieve
// Param 1: zipWriter is used for adding file to zip
// Param 2: filename is the target file to be added
func AddFileToZip(zipWriter *zip.Writer, filename string, fileToZip *os.File) error {
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// Using FileInfoHeader() above only uses the basename of the file. If we want
	// to preserve the folder structure we can overwrite this with the full path.
	header.Name = filename

	// Change to deflate to gain better compression
	// see http://golang.org/pkg/archive/zip/#pkg-constants
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
