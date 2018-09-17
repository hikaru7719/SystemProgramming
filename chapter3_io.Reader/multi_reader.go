package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	header := bytes.NewBufferString("--------Header--------\n")
	content := bytes.NewBufferString("Example of ip.MultiReader\n")
	footer := bytes.NewBufferString("--------Footer--------\n")

	reader := io.MultiReader(header, content, footer)
	io.Copy(os.Stdout, reader)
}
