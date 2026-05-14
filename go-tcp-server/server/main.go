package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {

	defer conn.Close()

	fmt.Println("New client connected:", conn.RemoteAddr())

	buffer := make([]byte, 1024)

	for {

		fmt.Println("Waiting for data from client...")
		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Println("client disconnected:", conn.RemoteAddr())
			return
		}
		fmt.Println("new client connected:", conn.RemoteAddr())
		fmt.Println("total bytes received:", n)
		fmt.Println("raw buffer: ", buffer[:n])

		fmt.Printf("Received data from client: %s\n", string(buffer[:n]))
	}
}

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

		go handleConnection(conn)
	}

}
