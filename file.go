package slitu

import (
	"bufio"
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
