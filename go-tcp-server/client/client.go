package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	defer conn.Close()

	message := "One piece is the best anime ever!"

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending message to server:", err)
		return
	}

	fmt.Println("Connected to server:", conn.RemoteAddr())
}
