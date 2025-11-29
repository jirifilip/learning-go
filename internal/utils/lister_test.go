package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func makeDirWithTempFiles(t *testing.T, dirname string, filenames []string) string {
	directory, err := os.MkdirTemp("", dirname)

	if err != nil {
		t.Fatal("Unable to create temporary directory")
	}

	for _, filename := range filenames {
		fullpath := filepath.Join(directory, filename)

		os.WriteFile(fullpath, []byte("Hello world!"), 0666)
	}

	return directory
}

func TestList(t *testing.T) {
	result := List()

	if result != 5 {
		t.Error("Result not equal to 5.")
	}
}

func TestGetFilenames(t *testing.T) {
	testFilenames := []string{"file1", "file2", "file3"}
	dirPath := makeDirWithTempFiles(t, "testingdir", testFilenames)

	result := GetFilenames(dirPath)

	if !slices.Equal(result, testFilenames) {
		t.Fatalf("Files not equal: %v != %v", result, testFilenames)
	}
}

func TestGetFiles(t *testing.T) {
	filenames := []string{"file1", "file22"}
	dirPath := makeDirWithTempFiles(t, "testingdir", filenames)

	infos := GetFiles(dirPath)

	fmt.Printf("%v\n", infos)
}
