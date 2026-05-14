package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Println("connected to server")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("enter message: ")

		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading input:", err)
			return
		}

		n, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("error writing:", err)
			return
		}

		fmt.Println("bytes sent:", n)
	}
}
