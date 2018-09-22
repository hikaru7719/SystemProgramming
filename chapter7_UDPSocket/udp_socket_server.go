package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Server is running at localhost:8888")
	conn, err := net.ListenPacket("udp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1500)
	for {
		length, remoteAddres, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Received from %v: %v\n", remoteAddres, string(buffer[:length]))
		_, err = conn.WriteTo([]byte("hello from server"), remoteAddres)
		if err != nil {
			panic(err)
		}
	}

}
