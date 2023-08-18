package slitu

import (
	"math/rand"
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

func TestCountFileLines(t *testing.T) {

	randLen := rand.Intn(10240)

	file, err := os.OpenFile("testfile", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		t.Fatalf("Failed to create test file: %s\n", err)
	}
	defer os.Remove("testfile")
	defer file.Close()

	for i := 0; i < randLen; i++ {
		file.Write([]byte{'a', '\n'})
	}

	n, err := CountFileLines("testfile")
	if err != nil {
		t.Fatalf("Failed to count: %s\n", err)
	}

	if n != randLen {
		t.Fatalf("Invalid result: want: %d, got: %d\n", randLen, n)
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
