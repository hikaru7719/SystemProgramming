package main

import (
	"os"
	"io"
)

func main(){
	file,err := os.Open("file_input.go")
	if err != nil{
		panic(err)
	}
	defer file.Close()
	//io.Copy(writer,reader)はreaderの内容をwriterに受け渡す。下記の例ではfileの内容を標準出力に受け渡している
	//io.Copy(file,file)でファイルのコピーができる。
	io.Copy(os.Stdout,file)
}
