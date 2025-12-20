# Directory Media File Comparator

This Go program compares the media files in two specified directories and identifies which files are unique to each directory.

## Features

- Compares media files (images and videos) between two directories.
- Normalizes filenames to handle case-insensitivity and different extensions for the same media type (e.g., `.jpeg` and `.jpg`).
- Generates output files listing the files that are missing from each of the directories.

## How to Use

1. **Configure Directories**: The directories to be compared are currently hardcoded in the `main` function in `main.go`.

    ```go
    folder1 := "./test/source/202501"
    folder2 := "./test/target/202501"
    ```

    You will need to modify these paths to the directories you wish to compare.

2. **Run the Program**: Execute the program from your terminal:

    ```sh
    go run main.go
    ```

## File Normalization

To ensure accurate comparison, the program normalizes filenames by:

- Converting the filename (without the extension) to lowercase.
- Standardizing media file extensions. For example, `.heic`, `.jpeg`, and `.jpg` are all treated as `.jpg`.

This prevents the program from incorrectly identifying files as different due to case variations or alternate extensions.

## Output

The program generates two files:

- `missing_from_folder1.txt`: Lists files that are present in the second directory but not in the first one. The output is a series of `move` commands.
- `missing_from_folder2.txt`: Lists files that are present in the first directory but not in the second one. The output is a series of `move` commands.
