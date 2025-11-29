package utils

import (
	"testing"
)

func TestList(t *testing.T) {
	result := List()

	if result != 5 {
		t.Error("Result not equal to 5.")
	}
}
