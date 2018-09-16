package main

import "os"

func main(){
	file, err := os.Create("test.txt")
	if err != nil{
		panic(err)
	}
	file.Write([]byte("os file example\n"))
	file.Close()
}
