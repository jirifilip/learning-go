package utils

import (
	"log"
	"os"
)

type FileInfo struct {
	Name  string
	IsDir bool
	Owner string
	Group string
}

func List() int {
	return 5
}

func getEntries(path string) []os.DirEntry {
	entries, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	return entries
}

func GetFilenames(path string) []string {
	entries := getEntries(path)

	filenames := make([]string, len(entries))

	for i, entry := range entries {
		filenames[i] = entry.Name()
	}

	return filenames
}

func GetFiles(path string) []FileInfo {
	entries := getEntries(path)

	infos := make([]FileInfo, len(entries))

	for i, entry := range entries {
		infos[i] = FileInfo{Name: entry.Name(), IsDir: entry.IsDir(), Owner: "unknown", Group: "unknown"}
	}

	return infos
}
