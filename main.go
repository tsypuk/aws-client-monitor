package main

import (
	"aws-client-monitor/docs"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net"
	"net/http"
	"sync"
	"time"
)

var broadcastChan = make(chan []byte)
var ch2 = make(chan []byte)

type ApiCallAttempt struct {
	Version        int    `json:"Version"`
	ClientId       string `json:"ClientId"`
	Type           string `json:"Type"`
	Service        string `json:"Service"`
	Api            string `json:"Api"`
	Timestamp      int64  `json:"Timestamp"`
	AttemptLatency int    `json:"AttemptLatency"`
	Fqdn           string `json:"Fqdn"`
	UserAgent      string `json:"UserAgent"`
	AccessKey      string `json:"AccessKey"`
	Region         string `json:"Region"`
	SessionToken   string `json:"SessionToken"`
	HttpStatusCode int    `json:"HttpStatusCode"`
	XAmzRequestId  string `json:"XAmzRequestId"`
	XAmzId2        string `json:"XAmzId2"`
}

// Struct for the ApiCall message
type ApiCall struct {
	Version             int    `json:"Version"`
	ClientId            string `json:"ClientId"`
	Type                string `json:"Type"`
	Service             string `json:"Service"`
	Api                 string `json:"Api"`
	Timestamp           int64  `json:"Timestamp"`
	AttemptCount        int    `json:"AttemptCount"`
	Region              string `json:"Region"`
	UserAgent           string `json:"UserAgent"`
	FinalHttpStatusCode int    `json:"FinalHttpStatusCode"`
	Latency             int    `json:"Latency"`
	MaxRetriesExceeded  int    `json:"MaxRetriesExceeded"`
}

// WebSocket upgrader
var upgrader = websocket.Upgrader{}

var clients = make(map[*websocket.Conn]bool)
var clientsLock sync.Mutex

func listenUDP(port int, ch []chan<- []byte) {
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
		for i := 0; i < len(ch); i++ {
			ch[i] <- buffer[:n]
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

// Goroutine-safe function to broadcast to all WebSocket clients
func broadcastMessages() {
	for {
		message := <-broadcastChan

		clientsLock.Lock()
		for client := range clients {

			var apiCall ApiCall
			err := json.Unmarshal(message, &apiCall)
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
					delete(clients, client)
				}
				continue
			}

			var apiCallAttempt ApiCallAttempt
			err = json.Unmarshal(message, &apiCallAttempt)
			if err != nil {
				print("Error unmarshalling ApiCall: %v", err)
			} else {
				fmt.Printf("Parsed ApiCall: %+v\n", apiCallAttempt)
				continue
			}

			print("Unknown message Type: %s", message)
		}
		clientsLock.Unlock()

	}
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} running
// @Router /status [get]
func statusHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "running",
	})
}

// WebSocket handler, to handle new connections
func wsHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: ", err)
		return
	}
	defer conn.Close()

	clientsLock.Lock()
	clients[conn] = true
	clientsLock.Unlock()

	defer func() {
		clientsLock.Lock()
		delete(clients, conn)
		clientsLock.Unlock()
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

// Serve the dashboard
func serveDashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", nil)
}

func main() {
	// Goroutine to listen on UDP and write to the channel
	go listenUDP(31000, []chan<- []byte{broadcastChan})

	// Goroutines to read from the channel
	//go writeToConsole(ch2)

	// Goroutine to broadcast the UDP data to WebSocket clients
	go broadcastMessages()

	// start web-server
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		v1.GET("/status", statusHandler)
		//eg := v1.Group("/example")
		//{
		//	eg.GET("/helloworld", Helloworld)
		//}
	}
	router.LoadHTMLFiles("templates/dashboard.html")

	// Serve WebSocket for live updates
	router.GET("/ws", wsHandler)

	// Serve the dashboard UI
	router.GET("/", serveDashboard)

	router.Static("/css", "./css")

	// Route to access the Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8080")

	// Prevent the main function from exiting
	for {
		time.Sleep(1 * time.Second)
	}
}
