package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	var buf bytes.Buffer

	run(&buf)

	actual := buf.String()
	expected := "Hello world! Listing 5\n"

	if actual != expected {
		t.Fatalf("Failed: '%s' != '%s'", actual, expected)
	}
}
