package main

import "os"

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("System Call Example\n"))
}
