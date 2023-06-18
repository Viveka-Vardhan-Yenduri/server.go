package main

import (
	"fmt"
	"net"
)

func handleUDPConnection(conn *net.UDPConn) {
	buffer := make([]byte, 1024)
	n, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	message := string(buffer[:n])
	fmt.Println("Received message:", message)

	// Sending a response back to the client
	response := []byte("Hello from UDP server!")
	_, err = conn.WriteToUDP(response, addr)
	if err != nil {
		fmt.Println("Error sending response:", err)
		return
	}
}

func main() {
	addr := &net.UDPAddr{
		Port: 8080,
		IP:   net.ParseIP("0.0.0.0"),
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error starting UDP server:", err)
		return
	}

	fmt.Println("UDP server started, waiting for connections...")

	for {
		handleUDPConnection(conn)
	}
}
