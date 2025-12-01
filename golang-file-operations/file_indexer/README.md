# Not-Google File Indexer

This Go program traverses a structured photo directory (Year -\> YearMonth), specifically looking for folders named `not_google`. It collects the full file paths of all files found within these specific folders and exports them to a text file.

## 📂 Expected Directory Structure

The program relies on a specific folder hierarchy to work correctly. It looks for:

1. **Year Folders** (The walker enters these)
2. **YearMonth Folders** (6 digits, e.g., `202501`)
3. **Target Folder** (Must be named `not_google`)

```text
Root Directory (e.g., Photos)
├── 2024/
│   ├── 202401/
│   │   ├── not_google/         <-- 🎯 Scans files inside here
│   │   │   ├── image01.jpg
│   │   │   └── video01.mp4
│   │   └── other_folder/       <-- Ignored
│   └── 202402/
├── 2025/
│   ├── 202511/
│   │   └── not_google/         <-- 🎯 Scans files inside here
│   └── ...
└── ...
```

## 🚀 Usage

### 1\. Prerequisites

- [Go (Golang)](https://go.dev/dl/) installed on your system.

### 2\. Configuration

The root directory path is currently hardcoded. You **must** update this to match your system before running.

### 3\. Running the Script

Open your terminal in the project directory and run:

```bash
go run main.go
```

Or build it into an executable:

```bash
go build -o indexer.exe
./indexer.exe
```

## 📄 Output

After execution, a file named **`all_not_google_files.txt`** will be created in the same directory as the script.

It contains the absolute paths of every file found:

```text
test\202501\not_google\GREEN ARROW.png
```

## 🧠 Technical Details

- **Filter Logic:** The script explicitly checks for folder names that are exactly 6 characters long and consist only of digits (to identify `YYYYMM` folders).
- **Optimization:** Once it finds a `YYYYMM` folder, it checks for `not_google`, scans it, and then skips the rest of the subdirectories in that month folder to save time.
