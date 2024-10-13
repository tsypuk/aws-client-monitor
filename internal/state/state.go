package state

import (
	"github.com/gorilla/websocket"
	"sync"
)

var BroadcastChan = make(chan []byte)
var Ch2 = make(chan []byte)
var Clients = make(map[*websocket.Conn]bool)
var ClientsLock sync.Mutex
