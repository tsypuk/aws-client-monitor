package handler

import (
	"aws-client-monitor/internal/state"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	CheckOrigin:      func(r *http.Request) bool { return true },
	HandshakeTimeout: time.Duration(time.Second * 5),
}

// WebSocket handler, to handle new connections
func WsHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		fmt.Println("Failed to set websocket upgrade: ", err)
		return
	}
	defer conn.Close()

	state.ClientsLock.Lock()
	state.Clients[conn] = true
	state.ClientsLock.Unlock()

	defer func() {
		state.ClientsLock.Lock()
		delete(state.Clients, conn)
		state.ClientsLock.Unlock()
	}()

	// Keep the connection open
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading WebSocket message:", err)
			break
		}
	}
}

// WebSocket handler, to handle new connections
func WsApiCallHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		fmt.Println("Failed to set websocket upgrade: ", err)
		return
	}
	defer conn.Close()

	state.ApiCallClientsLock.Lock()
	state.ApiCallClients[conn] = true
	state.ApiCallClientsLock.Unlock()

	defer func() {
		state.ApiCallClientsLock.Lock()
		delete(state.ApiCallClients, conn)
		state.ApiCallClientsLock.Unlock()
	}()

	// Keep the connection open
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading WebSocket message:", err)
			break
		}
	}
}
