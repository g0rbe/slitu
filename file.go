package slitu

import (
	"bufio"
	"bytes"
	"io"
	"os"
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
