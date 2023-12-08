package slitu

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

// IsLineInFile check whether s is exactly a line in file in path.
// A line in the file does not contains the newline character ("\n").
// This function uses bufio.Scanner to able to handle large files.
func IsLineInFile(path, s string) (bool, error) {

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == s {
			return true, nil
		}
	}

	return false, scanner.Err()
}

// CountFileLines returns the number lines in file.
func CountFileLines(path string) (int, error) {

	// Source: https://stackoverflow.com/questions/24562942/golang-how-do-i-determine-the-number-of-lines-in-a-file-efficiently

	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := file.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

// IsExists returns whether path is exists.
//
// It will panic is error is other than "no such file or directory"!
func IsExists(path string) bool {

	// _, err := os.Stat(path)
	// if err != nil {

	// 	if errors.Is(err, fs.ErrNotExist) {
	// 		return false
	// 	}

	// 	panic(fmt.Sprintf("Failed to stat %s: %s", path, err))
	// }

	// return true

	err := unix.Faccessat(unix.AT_FDCWD, path, syscall.F_OK, 0)
	if err != nil {

		if errors.Is(err, unix.ENOENT) {
			return false
		}

		panic(fmt.Sprintf("faccessat failed for %s: %s", path, err))
	}

	return true

}

// IsDir returns whether path is a directory..
//
// It will panic is error is other than fs.ErrNotExist!
func IsDir(path string) bool {

	stat, err := os.Stat(path)
	if err != nil {

		if errors.Is(err, fs.ErrNotExist) {
			return false
		}

		panic(fmt.Sprintf("Failed to stat %s: %s", path, err))
	}

	return stat.IsDir()
}

// IsExistsAndDir returns whether path is exists and a directory.
//
// It will panic is error is other than fs.ErrNotExist!
func IsExistsAndDir(path string) bool {
	return IsExists(path) && IsDir(path)
}
