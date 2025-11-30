package kopr

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/jirifilip/learning-go/testutils"
)

func TestExecCat(t *testing.T) {
	path := testutils.MakeDirWithTempFiles(t, "test", []string{"file2"})
	var buf bytes.Buffer

	CatFile(filepath.Join(path, "file2"), &buf)

	fileContent := buf.String()

	if fileContent != "Hello world!" {
		t.Fatalf("Content not expected: %s", fileContent)
	}
}
