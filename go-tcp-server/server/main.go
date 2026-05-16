package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var clients = make(map[net.Conn]bool)
var mu sync.Mutex

func broadcast(message string) {
	mu.Lock()
	defer mu.Unlock()
	for client := range clients {

		// intentionally slowing iteration
		// to increase overlap between goroutines
		time.Sleep(50 * time.Millisecond)

		_, err := client.Write([]byte(message))
		if err != nil {
			fmt.Println("write error:", err)
		}
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// shared state write
	mu.Lock()
	clients[conn] = true
	mu.Unlock()

	fmt.Println("new client connected:", conn.RemoteAddr())

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {

			// shared state delete
			mu.Lock()
			delete(clients, conn)
			mu.Unlock()

			fmt.Println("client disconnected:", conn.RemoteAddr())
			return
		}

		message := fmt.Sprintf("[%s]: %s", conn.RemoteAddr(), string(buffer[:n]))

		fmt.Print(message)

		// shared state iteration
		broadcast(message)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("server listening on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		go handleConnection(conn)
	}
}
