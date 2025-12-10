# File Missing - Directory Comparison Tool

A Go utility that compares two directories and identifies missing files by comparing filenames (without extensions).

## Overview

This tool compares the contents of two directories and identifies:
- Files present in directory A but missing in directory B
- Files present in directory B but missing in directory A

The comparison is done based on file basenames (filename without extension), making it useful for finding missing files across directories with potentially different file formats.

## Features

- **Case-insensitive comparison**: File names are converted to lowercase for comparison
- **Extension-agnostic**: Compares files by base name, ignoring extensions (e.g., `Image1.jpg` and `Image1.png` are considered the same)
- **Ignores directories**: Only compares files, nested directories are skipped
- **Exports results**: Generates two output files with lists of missing files from each directory

## Usage

### Configuration

Update the hardcoded directory paths in `main.go`:

```go
const dirA = "source" // Change as needed
const dirB = "target"                    // Change as needed
```

### Running the Program

```bash
go run main.go
```

### Output

The program generates two output files:
- **`missing_in_dir_B.txt`**: Contains files found in directory A but missing in directory B
- **`missing_in_dir_A.txt`**: Contains files found in directory B but missing in directory A

## How It Works

1. **Read directories**: Scans both directories and extracts file basenames
2. **Build maps**: Creates maps of file basenames for efficient comparison
3. **Find differences**: Identifies files in each directory that don't exist in the other
4. **Export results**: Writes the missing file lists to output text files

## Example

Given two directories:

**Directory A:**
- Image1.jpg
- Video2.mp4
- Document.pdf

**Directory B:**
- Image1.png
- Document.docx
- Audio3.wav

**Results:**

`missing_in_dir_B.txt`:
```
Video2
```

`missing_in_dir_A.txt`:
```
Audio3
```

Note: `Image1` and `Document` are not listed as missing because they exist in both directories (even though with different extensions).

## Functions

- **`getFileBases(dirPath string)`**: Reads a directory and returns a map of file basenames
- **`findMissing(source, target map[string]bool)`**: Compares two file maps and returns missing files
- **`exportResults(missing []string, outputFileName string, sourceDir string)`**: Writes results to a file
