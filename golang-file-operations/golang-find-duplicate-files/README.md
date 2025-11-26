# 📂 Duplicate File Finder (YYYY/YYYYMM Structure)

A Python script to identify and list files duplicated between monthly directories (`YYYYMM`) and a specific subdirectory named `not_google` within the same month, storing the results in a text file.

---

## ✨ Features

- **Recursive Scanning:** Iterates through all year folders (`YYYY`) and month folders (`YYYYMM`).
- **Duplicate Detection:** Compares files within the `YYYYMM` directory against files in the `not_google` subdirectory.
- **Detailed Output:** Generates a text file listing the **full paths** of all detected duplicate files.
- **Simple Usage:** Requires only specifying the base directory to begin the search.
