package utils

import (
	"log"
	"os"
)

func List() int {
	return 5
}

func GetFilenames(path string) []string {
	entries, err := os.ReadDir(path)

	if err != nil {
		log.Panic("Could not list directory!")
	}

	filenames := make([]string, len(entries))

	for i, entry := range entries {
		filenames[i] = entry.Name()
	}

	return filenames
}
