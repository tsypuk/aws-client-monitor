package handler

import (
	"aws-client-monitor/internal/state"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{}

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