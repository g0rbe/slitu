package slitu

import "testing"

func TestSHA256Str(t *testing.T) {

	v := SHA256Str("test")
	if v != "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08" {
		t.Fatalf("FAIL: string mismatch (%s)", v)
	}
}

func BenchmarkSHA256Str(b *testing.B) {

	for i := 0; i < b.N; i++ {
		SHA256Str("test")
	}
}
