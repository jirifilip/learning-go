package utils

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestList(t *testing.T) {
	result := List()

	if result != 5 {
		t.Error("Result not equal to 5.")
	}
}

func TestGetFilenames(t *testing.T) {
	directory, err := os.MkdirTemp("", "testdir")

	if err != nil {
		t.Fatal("Unable to create temporary directory")
	}

	testFilenames := []string{"file1", "file2", "file99"}

	for _, filename := range testFilenames {
		fullpath := filepath.Join(directory, filename)

		os.WriteFile(fullpath, []byte("Hello world!"), 0666)
	}

	result := GetFilenames(directory)

	if !slices.Equal(result, testFilenames) {
		t.Fatalf("Files not equal: %v != %v", result, testFilenames)
	}
}
