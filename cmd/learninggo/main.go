package main

import (
	"fmt"
	"io"
	"os"

	"github.com/jirifilip/learning-go/internal/utils"
)

func main() {
	run(os.Stdout)
}

func run(writer io.Writer) {
	message := fmt.Sprintf("Hello world! Listing %d\n", utils.List())

	writer.Write([]byte(message))
}
