package testutils

import (
	"os"
	"path/filepath"
	"testing"
)

func MakeDirWithTempFiles(t *testing.T, dirname string, filenames []string) string {
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
