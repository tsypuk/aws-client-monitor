package state

import (
	"aws-client-monitor/internal/domain"
	"github.com/gorilla/websocket"
	"sync"
)

var BroadcastChan = make(chan domain.UdpPayload)
var LoggingChan = make(chan domain.UdpPayload)
var Clients = make(map[*websocket.Conn]bool)
var ClientsLock sync.Mutex
