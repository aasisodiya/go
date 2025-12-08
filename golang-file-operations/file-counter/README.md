# Go Recursive Directory File Counter

A lightweight Go utility designed to traverse a directory tree recursively. It provides a detailed file count for the root directory and every individual subfolder, along with a grand total of all files found.

## Features

* **Recursive Traversal:** Walks through the specified root directory and all nested subdirectories.
* **Granular Reporting:** Prints the file count for *each* specific folder found in the tree.
* **System File Exclusion:** Automatically ignores `desktop.ini` files to ensure accurate counts of user data (useful for Windows/Synology backups).
* **Sorted Output:** displays folder paths alphabetically for better readability.
* **Total Aggregation:** Calculates and displays the sum of all files across the entire directory structure.

## Configuration

Currently, the target directory is hardcoded in the `main` function. Before running the script, you must update the path to match your target folder.

1.  Open `main.go`.
2.  Locate **Line 68**:
    ```go
	targetDir := "./test/" // Hardcoded path
    ```
3.  Change the string to the directory you wish to scan.

## Usage

### Prerequisites
* [Go](https://go.dev/dl/) installed on your machine.

### Running the Script

1.  Navigate to the directory containing your file (e.g., `main.go`).
2.  Run the following command in your terminal:

```bash
go run main.go
```

## Example Output

```text
--- File Count Per Folder (in C:\Users\akash\Documents\github\go\golang-file-operations\file-counter\test) ---
Root Directory (C:\Users\akash\Documents\github\go\golang-file-operations\file-counter\test): 0 files
Subfolder: 202501: 2 files
Subfolder: 202501\not_google: 1 files
Subfolder: 202502: 2 files
Subfolder: 202502\temp: 1 files
-----------------------------------------------------
Total files in all folders and subfolders: 6
```

## Future Improvements

The code contains commented-out logic to accept command-line arguments. To enable dynamic path selection without editing the code:

- Uncomment lines 64-67 in `main.go`.
- Update line 68 to use `os.Args[1]` instead of the hardcoded string.
- Run via: `go run main.go "C:\Path\To\Folder"`