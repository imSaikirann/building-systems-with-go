package main

import (
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	fmt.Println("Server is listening on port 8080...")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("New client connected:", conn.RemoteAddr())

		conn.Close()
	}

}
