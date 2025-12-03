# Go Directory Comparator

A simple, lightweight Go tool to recursively compare the contents of two directories. It checks for file existence and file size differences, useful for verifying backups or syncing folders.

## Features

- **Recursive Traversal:** Scans all subdirectories within the specified root folders.
- **Case-Insensitive Matching:** Normalizes file paths to lowercase before comparing (useful for Windows environments).
- **Size Verification:** Flags files that exist in both locations but have different file sizes.
- **Report Generation:** Generates two text files detailing discrepancies.

## Prerequisites

- [Go](https://go.dev/dl/) installed on your machine.

## Configuration

**Note:** Currently, the directory paths are hardcoded in the source file. You must update them before running the program.

1. Open `main.go`.
2. Locate the `main` function (lines 14-15).
3. Update `dir1` and `dir2` variables with the absolute paths you want to compare:

```go
func main() {
    dir1 := "C:/path/to/source/folder"      // Update this path
    dir2 := "C:/path/to/destination/folder" // Update this path
    // ...
}
```

## Output

The program creates two output files in the current working directory:

1. only_in_dir1.txt

   - Lists files present in dir1 but missing in dir2.
   - Also lists files present in both but with different sizes (relative to dir1).

2. only_in_dir2.txt
   - Lists files present in dir2 but missing in dir1.
   - Also lists files present in both but with different sizes (relative to dir2).

## Output Example

```text
202501\batman.png
202502\cyborg.png
```

## How it Works

1. `Collection`: The program walks through both directories.
2. `Indexing`: It creates a map for each directory where the key is the lowercased relative path and the value is the file size.
3. `Comparison`: It iterates through the maps to find missing keys (files) or mismatched values (sizes).
