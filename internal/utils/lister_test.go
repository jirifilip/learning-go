package utils

import (
	"fmt"
	"slices"
	"testing"

	"github.com/jirifilip/learning-go/testutils"
)

func TestList(t *testing.T) {
	result := List()

	if result != 5 {
		t.Error("Result not equal to 5.")
	}
}

func TestGetFilenames(t *testing.T) {
	testFilenames := []string{"file1", "file2", "file3"}
	dirPath := testutils.MakeDirWithTempFiles(t, "testingdir", testFilenames)

	result := GetFilenames(dirPath)

	if !slices.Equal(result, testFilenames) {
		t.Fatalf("Files not equal: %v != %v", result, testFilenames)
	}
}

func TestGetFiles(t *testing.T) {
	filenames := []string{"file1", "file22"}
	dirPath := testutils.MakeDirWithTempFiles(t, "testingdir", filenames)

	infos := GetFiles(dirPath)

	fmt.Printf("%v\n", infos)
}
