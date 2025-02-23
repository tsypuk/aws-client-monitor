package main

import (
	"aws-client-monitor/internal/domain"
	"aws-client-monitor/internal/router"
	"aws-client-monitor/internal/state"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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

func writeToConsole(ch <-chan domain.UdpPayload) {
	for msg := range ch {
		if apiType, err := domain.NewApiBaseType(msg); err == nil {
			switch apiType.Type {
			case "ApiCall":
				if apiCall, err := domain.NewApiCall(msg); err == nil {
					if err := apiCall.Validate(); err == nil {
						t := time.Unix(apiCall.Timestamp/1000, 0).UTC()
						humanReadable := t.Format("2006-01-02 15:04:05 UTC")
						fmt.Printf("%s [Client:%s(%s) ==========> AWS:%s:%s] : %dms %s Code:%d\n",
							humanReadable, apiCall.ClientId, apiCall.UserAgent[:25], apiCall.Service, apiCall.Api, apiCall.Latency, apiCall.Region, apiCall.FinalHttpStatusCode)
						if apiCall.FinalAwsException != "" {
							fmt.Printf("Error: %s (%s)\n", apiCall.FinalAwsException, apiCall.FinalAwsExceptionMessage)
						}
					} else {
						fmt.Println("Error event validation")
					}
				}

			case "ApiCallAttempt":
				if apiCallAttempt, err := domain.NewApiCallAttempt(msg); err == nil {
					if err := apiCallAttempt.Validate(); err == nil {
						t := time.Unix(apiCallAttempt.Timestamp/1000, 0).UTC()
						humanReadable := t.Format("2006-01-02 15:04:05 UTC")
						fmt.Printf("%s [Client:%s(%s) ==========> AWS:%s:%s(%s)] : %dms %s Code:%d\n",
							humanReadable, apiCallAttempt.ClientId, apiCallAttempt.UserAgent[:25], apiCallAttempt.Service, apiCallAttempt.Api, apiCallAttempt.Fqdn, apiCallAttempt.AttemptLatency, apiCallAttempt.Region, apiCallAttempt.HttpStatusCode)
						if apiCallAttempt.AwsException != "" {
							fmt.Printf("Error: %s (%s)\n", apiCallAttempt.AwsException, apiCallAttempt.AwsExceptionMessage)
						}
					} else {
						fmt.Println("Error event validation")
					}
				}
			}
		}
	}
}

// Channel slice to hold active WebSocket connections
//var wsClients = make([]*websocket.Conn, 0)

// Goroutine-safe function to broadcast to all WebSocket Clients
func broadcastMessages() {
	for {
		message := <-state.BroadcastChan

		// ApiCall Websocket
		state.ApiCallClientsLock.Lock()
		for client := range state.ApiCallClients {

			if apiType, err := domain.NewApiBaseType(message); err == nil {
				switch apiType.Type {
				case "ApiCall":
					if apiCall, err := domain.NewApiCall(message); err == nil {
						if err := apiCall.Validate(); err == nil {
							if err = client.WriteJSON(apiCall); err != nil {
								// fmt.Println("Error sending WebSocket message:", err)
								client.Close() // Close the connection if there's an error
								delete(state.ApiCallClients, client)
							}
						}
					}

				case "ApiCallAttempt":
					if apiCallAttempt, err := domain.NewApiCallAttempt(message); err == nil {
						if err := apiCallAttempt.Validate(); err == nil {
							if err = client.WriteJSON(apiCallAttempt); err != nil {
								// fmt.Println("Error sending WebSocket message:", err)
								client.Close() // Close the connection if there's an error
								delete(state.ApiCallClients, client)
							}
						}
					}
				}
			} else {
				print(fmt.Sprintf("error extracting API Type from message: %v, got error: %v", message, err))
			}
		}
		state.ApiCallClientsLock.Unlock()
	}
}

func main() {
	// Goroutine to listen on UDP and write to the channel
	go listenUDP(31000, []chan<- domain.UdpPayload{state.BroadcastChan, state.LoggingChan})
	//go listenUDP(31000, []chan<- domain.UdpPayload{state.BroadcastChan})

	// Goroutines to read from the channel
	go writeToConsole(state.LoggingChan)

	// Goroutine to broadcast the UDP data to WebSocket Clients
	go broadcastMessages()

	// start web-server
	if err := router.CreateRouter(gin.Default()).Run(":8080"); err != nil {
		log.Fatal(err)
	}

	// Prevent the main function from exiting
	for {
		time.Sleep(1 * time.Second)
	}
}
