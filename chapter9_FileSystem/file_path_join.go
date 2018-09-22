package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Printf("Temp FIle Path: %s\n", filepath.Join(os.TempDir(), "text.txt"))
}
