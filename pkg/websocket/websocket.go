package websocket

import (
	"annotation-service/pkg/config"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var idMutex sync.Mutex
var clientID int

// Generates a unique client ID
func GenerateClientID() int {
	idMutex.Lock()
	defer idMutex.Unlock()
	clientID++
	return clientID
}

// Resets the client ID counter
func ResetClientID() {
	idMutex.Lock()
	defer idMutex.Unlock()
	clientID = 0
}

// Returns a configured WebSocket upgrader
func GetUpgrader() websocket.Upgrader {
	return websocket.Upgrader{
		ReadBufferSize:  config.AppConfig.WebSocket.ReadBufferSize,
		WriteBufferSize: config.AppConfig.WebSocket.WriteBufferSize,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}
