# Directory Structure Comparator

A simple Go utility to recursively compare the folder structures of two different directories. It identifies missing subdirectories in both locations and outputs the differences to text files.

## Features

- **Recursive Traversal:** Scans all subdirectories within the specified root paths.
- **Folder-Only Comparison:** Ignores files and focuses strictly on directory hierarchy.
- **Case-Insensitive:** Compares relative paths in lowercase (e.g., `Folder/A` matches `folder/a`).
- **Dual Output:** Generates separate reports for unique folders in Directory 1 and Directory 2.

## Prerequisites

- [Go (Golang)](https://go.dev/dl/) installed on your machine.

## Configuration

Before running the script, you must define the two directories you wish to compare.

1. Open `main.go`.
2. Locate the `main()` function.
3. Update the `dir1` and `dir2` variables with your specific paths:

```go
func main() {
    // Update these paths
    dir1 := "C:/Path/To/Your/First/Directory"
    dir2 := "C:/Path/To/Your/Second/Directory"

    // ... rest of the code
}
```

## Logic Details

- The script creates a map of the relative paths from both root directories.
- It normalizes these relative paths to lowercase to ensure the comparison is case-insensitive.
- If a matching relative path (key) is not found in the opposing map, the folder is flagged as unique.
