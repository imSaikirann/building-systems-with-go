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

		buffer := make([]byte, 1024)

		for {
			n, err := conn.Read(buffer)

			if err != nil {
				fmt.Println("Error reading from client:", err)
				return
			}
			fmt.Println("new client connected:", conn.RemoteAddr())
			fmt.Println("total bytes received:", n)
			fmt.Println("raw buffer: ", buffer[:n])

			fmt.Printf("Received data from client: %s\n", string(buffer[:n]))
		}

	}

}
