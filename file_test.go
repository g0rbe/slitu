package slitu

import (
	"os"
	"testing"
)

func TestIsLineInFile(t *testing.T) {

	file, err := os.OpenFile("testfile", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		t.Fatalf("Failed to create testfile: %s\n", err)
	}
	defer os.Remove("testfile")
	defer file.Close()

	if _, err := file.WriteString("a\nb\nc\nd\n"); err != nil {
		t.Fatalf("Failed to write to testfile: %s\n", err)
	}

	exist, err := IsLineInFile("testfile", "d")
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	if !exist {
		t.Fatalf("\"d\" should exist in testfile\n")
	}
}

func BenchmarkIsLineinFile(b *testing.B) {

	file, err := os.OpenFile("testfile", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		b.Fatalf("Failed to create testfile: %s\n", err)
	}
	defer os.Remove("testfile")
	defer file.Close()

	if _, err := file.WriteString("a\nb\nc\nd\n"); err != nil {
		b.Fatalf("Failed to write to testfile: %s\n", err)
	}

	file.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IsLineInFile("testfile", "d")
	}

}
