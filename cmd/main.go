package main

import (
	"aws-client-monitor/internal/domain"
	"aws-client-monitor/internal/router"
	"aws-client-monitor/internal/state"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net"
	"time"
)

func listenUDP(port int, ch []chan<- domain.UdpPayload) {
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
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}
		udpPayload := domain.UdpPayload{
			UDPAddr: addr,
			Payload: buffer[:n],
		}
		// Send received data to channel
		for i := 0; i < len(ch); i++ {
			//ch[i] <- buffer[:n]
			ch[i] <- udpPayload
		}
	}
}

func writeToConsole(ch <-chan []byte) {
	for msg := range ch {
		fmt.Println("Received from channel:", string(msg))
	}
}

// Channel slice to hold active WebSocket connections
//var wsClients = make([]*websocket.Conn, 0)

// Goroutine-safe function to broadcast to all WebSocket Clients
func broadcastMessages() {
	for {
		message := <-state.BroadcastChan

		print(message.UDPAddr.IP.String())

		state.ClientsLock.Lock()
		for client := range state.Clients {

			var apiCall domain.ApiCall
			err := json.Unmarshal(message.Payload, &apiCall)
			if err != nil {
				print("Error unmarshalling ApiCall: %v", err)
			} else {
				fmt.Printf("Parsed ApiCall: %+v\n", apiCall)
				seconds := apiCall.Timestamp / 1000
				nanoseconds := (apiCall.Timestamp % 1000) * 1_000_000

				// Convert Unix timestamp to time.Time
				dt := time.Unix(seconds, nanoseconds)
				// Generate a random color in hex format
				color := "#FF0000"

				if apiCall.FinalHttpStatusCode == 200 {
					color = "00FF00"
				}

				//dt.Format(time.RFC3339),
				err = client.WriteJSON(map[string]interface{}{
					"datetime": dt.Format("2006-01-02 15:04:05.000"),
					"latency":  apiCall.Latency,
					"color":    color,
					"api":      fmt.Sprintf("%s:%s", apiCall.Service, apiCall.Api),
					"service":  apiCall.Service,
					"response": apiCall.FinalHttpStatusCode,
				})
				if err != nil {
					fmt.Println("Error sending WebSocket message:", err)
					client.Close() // Close the connection if there's an error
					delete(state.Clients, client)
				}
				continue
			}

			var apiCallAttempt domain.ApiCallAttempt
			err = json.Unmarshal(message.Payload, &apiCallAttempt)
			if err != nil {
				print("Error unmarshalling ApiCall: %v", err)
			} else {
				fmt.Printf("Parsed ApiCall: %+v\n", apiCallAttempt)
				continue
			}

			print("Unknown message Type: %s", message.Payload)
		}
		state.ClientsLock.Unlock()
	}
}

func main() {
	// Goroutine to listen on UDP and write to the channel
	go listenUDP(31000, []chan<- domain.UdpPayload{state.BroadcastChan})

	// Goroutines to read from the channel
	//go writeToConsole(ch2)

	// Goroutine to broadcast the UDP data to WebSocket Clients
	go broadcastMessages()

	// start web-server
	rout := router.CreateRouter(gin.Default())
	rout.Run(":8080")

	// Prevent the main function from exiting
	for {
		time.Sleep(1 * time.Second)
	}
}