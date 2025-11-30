package kopr

import (
	"io"
	"log"
	"os/exec"
)

func CatFile(filename string, writer io.Writer) {
	command := exec.Command("cat", filename)

	command.Stdout = writer
	command.Stderr = writer

	err := command.Run()
	if err != nil {
		log.Fatal(err)
	}
}
