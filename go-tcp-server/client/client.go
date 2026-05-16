package main

import (
	"fmt"
	"net"
)

func readMessages(conn net.Conn) {
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("server disconnected:", err)
			return
		}

		fmt.Print(string(buffer[:n]))
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Println("connected to server")

	go readMessages(conn)

	// intentionally spamming messages
	// to create concurrent overlap
	for {
		_, err := conn.Write([]byte("spam message\n"))
		if err != nil {
			fmt.Println("write error:", err)
			return
		}
	}
}
