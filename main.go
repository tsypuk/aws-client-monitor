package main

import (
	"fmt"
	"net"
	"time"
)

func listenUDP(port int, ch chan<- []byte) {
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("0.0.0.0"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error listening on UDP:", err)
		return
	}
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			print("Error closing UDP connection:", err)
		}
	}(conn)

	buffer := make([]byte, 2048)
	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}
		// Send received data to channel
		ch <- buffer[:n]
	}
}

func writeToConsole(ch <-chan []byte) {
	for msg := range ch {
		fmt.Println("Received from channel:", string(msg))
	}
}

func main() {
	byteChannel := make(chan []byte)

	// Goroutine to listen on UDP and write to the channel
	go listenUDP(31000, byteChannel)

	// Goroutines to read from the channel
	go writeToConsole(byteChannel)

	// Prevent the main function from exiting
	for {
		time.Sleep(1 * time.Second)
	}
}
