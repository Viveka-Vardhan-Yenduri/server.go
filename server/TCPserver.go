package main

import (
	"fmt"
	"net"
)

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	message := string(buffer)
	fmt.Println("Received message:", message)

	// Sending a response back to the client
	response := "Hello from TCP server!"
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error sending response:", err)
		return
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}

	fmt.Println("TCP server started, waiting for connections...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleTCPConnection(conn)
	}
}
