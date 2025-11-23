# 📸 Media File Synchronization Checker

- [📸 Media File Synchronization Checker](#-media-file-synchronization-checker)
  - [**Project Goal**](#project-goal)
  - [**Features**](#features)
  - [**Folder Structure Assumption**](#folder-structure-assumption)
  - [**How the Comparison Works**](#how-the-comparison-works)

---

## **Project Goal**

This script is designed to compare the media files (photos and videos) contained within two separate large directory structures—typically a **Backup** and an **External Drive**—to identify any files that are **missing** in either location.

The comparison is intelligent, accounting for a specific nested structure and considering a file present even if it has moved between the primary monthly folder and a special subfolder within the same month across the two drives.

## **Features**

- **Deep, Nested Comparison:** Recursively compares files within a highly specific `yyyy/yyyymm` and `yyyy/yyyymm/not_google` folder structure.
- **Flexible Presence Check:** A file is considered **"present"** in a given month on a drive if it exists in _either_ the `yyyymm` folder or its `yyyymm/not_google` subfolder.
- **Comprehensive File Matching:** Compares all file types (photos, videos, documents, etc.) without skipping any formats.
- **Detailed Missing File Report:** Generates lists of files missing from the **Backup** and files missing from the **External Drive**.
- **Hardcoded Paths:** The paths to the two main directories are explicitly set within the script for consistent execution.

## **Folder Structure Assumption**

The script is built around the following directory hierarchy for both your **Backup** and **External Drive** roots:

```bash
[DRIVE_ROOT]
└── yyyy (Year folder, e.g., 2023)
    └── yyyymm (Month folder, e.g., 202312)
        ├── file1.jpg
        ├── file2.mp4
        └── not_google (Special subfolder)
            ├── file3.png
            └── file4.mov
```

The script will iterate through all `yyyy` folders and then all `yyyymm` folders found within them to perform the comparison.

## **How the Comparison Works**

For any given month (e.g., `202312`), the script performs the following logic:

1. **Collect All Files:** It compiles a list of _all_ files (using only the filename, not the path) from:

   - `[BACKUP_ROOT]/yyyy/yyyymm/*`
   - `[BACKUP_ROOT]/yyyy/yyyymm/not_google/*`

2. **Compare Sets:** It compares this combined list of files against the combined list of files from the **External Drive's** corresponding monthly location.
3. **Identify Missing:**
   - **Missing in Backup:** Files present in the External Drive's combined monthly set but **not** in the Backup's combined monthly set.
   - **Missing in External Drive:** Files present in the Backup's combined monthly set but **not** in the External Drive's combined monthly set.
