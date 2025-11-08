// Code to sync files between two directories based on filename and modification time. But first we need to compare files and display them that needs to be synced and then write those file name with path to a text file and also include the operation needed (copy from source to target or delete from target). Once this file is created. User will manually verify and if he is happy then he will run the code again with a flag to actually perform the operations mentioned in the text file.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Main entry point for the filesync command.
//
// Filesync is a command to help keep two directories in sync.
// It compares the files in the two directories and generates a plan
// to bring the target directory up to date with the source directory.
//
// The plan is written to a file, and can be verified by the user.
// If the user is happy with the plan, they can re-run the command
// with the -apply flag, which will perform the operations listed in
// the plan file.
//
// The command takes three flags: -source, -target, and -plan-file.
// -source is the path to the source directory, -target is the path to
// the target directory, and -plan-file is the path to the plan file.
//
// The command will exit with a non-zero status if either the source or
// target directories are not provided.
//
// If the -apply flag is provided, the command will read the plan file and
// perform the operations listed in it. If the plan file does not exist,
// the command will exit with a non-zero status.
//
// If the -apply flag is not provided, the command will generate a plan
// and write it to the plan file. If the plan file already exists, it
// will be overwritten. If the plan is empty (i.e. there are no operations
// required to bring the target directory up to date), the plan file will
// be deleted.
//
// The command will print the planned operations to stdout, and will also
// write them to the plan file. If the plan is empty, it will print a
// message to stdout indicating that no actions are required.
func main() {
	srcRoot := flag.String("source", "", "source directory")
	dstRoot := flag.String("target", "", "target directory")
	planPath := flag.String("plan-file", "filesync_plan.txt", "path to plan file")
	apply := flag.Bool("apply", false, "when set, read the plan file and perform the operations")
	flag.Parse()

	if *srcRoot == "" || *dstRoot == "" {
		fmt.Fprintln(os.Stderr, "source and target must be provided")
		flag.Usage()
		os.Exit(2)
	}

	absSrc, err := filepath.Abs(*srcRoot)
	if err != nil {
		fatal(err)
	}
	absDst, err := filepath.Abs(*dstRoot)
	if err != nil {
		fatal(err)
	}

	if *apply {
		if err := executePlan(*planPath); err != nil {
			fatal(err)
		}
		return
	}

	plan, err := createPlan(absSrc, absDst)
	if err != nil {
		fatal(err)
	}

	if len(plan) == 0 {
		fmt.Println("No actions required. Nothing to sync.")
		// ensure empty plan file exists
		_ = os.WriteFile(*planPath, []byte{}, 0644)
		return
	}

	fmt.Println("Planned operations:")
	for _, ln := range plan {
		fmt.Println(ln)
	}

	if err := os.WriteFile(*planPath, []byte(strings.Join(plan, "\n")+"\n"), 0644); err != nil {
		fatal(err)
	}
	fmt.Printf("\nPlan written to %s\n", *planPath)
	fmt.Println("Verify the plan file. Re-run with -apply to perform the listed operations.")
}

// createPlan compares files under srcRoot and dstRoot and returns plan lines.
// Lines are of form:
// COPY <absSourcePath> -> <absTargetPath>
// DELETE <absTargetPath>
//
// createPlan first walks the source directory and builds a map of
// relative paths to file info. Then it walks the target directory and
// builds a map of relative paths to file info. Finally, it compares the
// two maps and generates plan lines accordingly.
func createPlan(srcRoot, dstRoot string) ([]string, error) {
	srcMap := map[string]fs.FileInfo{}
	dstMap := map[string]fs.FileInfo{}

	// Walk source and build map of relative paths to file info
	if err := filepath.WalkDir(srcRoot, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(srcRoot, p)
		if err != nil {
			return err
		}
		// Use slash-separated relative path for comparison on all OSes
		rel = filepath.ToSlash(rel)
		srcMap[rel] = info
		return nil
	}); err != nil {
		return nil, err
	}

	// Walk target and build map of relative paths to file info
	if err := filepath.WalkDir(dstRoot, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(dstRoot, p)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		dstMap[rel] = info
		return nil
	}); err != nil {
		return nil, err
	}

	var plan []string
	// For each file in source: if missing in target => COPY. If present but source is newer => COPY.
	for rel, sinfo := range srcMap {
		dstInfo, exists := dstMap[rel]
		srcPath := filepath.Join(srcRoot, filepath.FromSlash(rel))
		dstPath := filepath.Join(dstRoot, filepath.FromSlash(rel))

		if !exists {
			plan = append(plan, fmt.Sprintf("COPY %s -> %s", srcPath, dstPath))
			continue
		}
		// If source modified strictly after target, plan copy. Use time comparison with a small epsilon to avoid minor FS differences.
		if sinfo.ModTime().After(dstInfo.ModTime().Add(1 * time.Second)) {
			plan = append(plan, fmt.Sprintf("COPY %s -> %s", srcPath, dstPath))
		}
	}

	// For each file in target but not in source => DELETE
	for rel := range dstMap {
		if _, exists := srcMap[rel]; !exists {
			dstPath := filepath.Join(dstRoot, filepath.FromSlash(rel))
			plan = append(plan, fmt.Sprintf("DELETE %s", dstPath))
		}
	}

	return plan, nil
}

// executePlan reads a plan file generated by planFiles and executes the operations specified in the plan.
// The plan file is expected to be in the following format:
//
// # Comment lines are ignored
// COPY <src> -> <dst>
// DELETE <dst>
//
// The function will return an error if any of the operations in the plan fail.
// If an error occurs, the function will continue to process the plan and return the last error encountered.
func executePlan(planPath string) error {
	f, err := os.Open(planPath)
	if err != nil {
		return fmt.Errorf("open plan file: %w", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var lastErr error
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		switch {
		case strings.HasPrefix(line, "COPY "):
			// format: COPY <src> -> <dst>
			parts := strings.SplitN(line[len("COPY "):], "->", 2)
			if len(parts) != 2 {
				fmt.Fprintf(os.Stderr, "invalid COPY line: %s\n", line)
				lastErr = fmt.Errorf("invalid COPY line")
				continue
			}
			src := strings.TrimSpace(parts[0])
			dst := strings.TrimSpace(parts[1])
			fmt.Printf("Copying: %s -> %s\n", src, dst)
			if err := ensureDirAndCopy(src, dst); err != nil {
				fmt.Fprintf(os.Stderr, "copy error: %v\n", err)
				lastErr = err
			}
		case strings.HasPrefix(line, "DELETE "):
			// format: DELETE <dst>
			dst := strings.TrimSpace(line[len("DELETE "):])
			// Create deleted folder in the same parent directory
			deletedDir := filepath.Join(filepath.Dir(dst), "deleted")
			if err := os.MkdirAll(deletedDir, 0o755); err != nil {
				fmt.Fprintf(os.Stderr, "create deleted dir error: %v\n", err)
				lastErr = err
				continue
			}
			// Move file to deleted folder
			newPath := filepath.Join(deletedDir, filepath.Base(dst))
			fmt.Printf("Moving to deleted folder: %s -> %s\n", dst, newPath)
			if err := os.Rename(dst, newPath); err != nil {
				if os.IsNotExist(err) {
					continue
				}
				fmt.Fprintf(os.Stderr, "move to deleted folder error: %v\n", err)
				lastErr = err
			}
		default:
			fmt.Fprintf(os.Stderr, "unknown plan line: %s\n", line)
			lastErr = fmt.Errorf("unknown plan line")
		}
	}
	if err := sc.Err(); err != nil {
		return err
	}
	return lastErr
}

// EnsureDirAndCopy ensures that the destination directory exists and
// copies the source file to the destination, preserving file permissions
// and modification time. If the destination file already exists, it will
// be overwritten. If an error occurs during copying, the temporary file
// will be removed and the original file will not be modified.
func ensureDirAndCopy(src, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return fmt.Errorf("source is a directory: %s", src)
	}

	dir := filepath.Dir(dst)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	// perform file copy
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	tmpDst := dst + ".tmp.copy"
	out, err := os.Create(tmpDst)
	if err != nil {
		return err
	}
	if _, err := io.Copy(out, in); err != nil {
		out.Close()
		_ = os.Remove(tmpDst)
		return err
	}
	if err := out.Sync(); err != nil {
		out.Close()
		_ = os.Remove(tmpDst)
		return err
	}
	if err := out.Close(); err != nil {
		_ = os.Remove(tmpDst)
		return err
	}
	// set permissions to match source
	if err := os.Chmod(tmpDst, info.Mode()); err != nil {
		_ = os.Remove(tmpDst)
		return err
	}
	// set modtime
	if err := os.Chtimes(tmpDst, time.Now(), info.ModTime()); err != nil {
		_ = os.Remove(tmpDst)
		return err
	}
	// atomic rename
	if err := os.Rename(tmpDst, dst); err != nil {
		_ = os.Remove(tmpDst)
		return err
	}
	return nil
}

// fatal prints an error to stderr and exits with status 1.
// It is intended for use in main functions to report errors that
// cannot be recovered from.
func fatal(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}
