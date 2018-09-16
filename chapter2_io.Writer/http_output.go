package main

import (
	"net"
	"io"
	"os"
)

func main(){
	conn,  err := net.Dial("tcp","ascii.jp:80")
	if err != nil{
		panic(err)
	}

	//内部的にはconn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")を呼び出している
	io.WriteString(conn,"GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")
	io.Copy(os.Stdout,conn)
}
