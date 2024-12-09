package chat

import (
	"annotation-service/pkg/database"
	"annotation-service/pkg/websocket"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type ChatMessage struct {
	ProjectID string `json:"project_id"`
	PartID    string `json:"part_id"`
	Message   string `json:"message"`
	Sender    string `json:"sender"`
	Recipient string `json:"to"`
	Timestamp string `json:"timestamp"`
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	clientID := websocket.GenerateClientID()
	log.Printf("New chat client connected with ID: %d\n", clientID)

	upgrader := websocket.GetUpgrader()
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	messageChan := make(chan ChatMessage)

	//Go routine
	go processChatMessages(clientID, messageChan)

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message from client %d: %v\n", clientID, err)
			break
		}

		log.Printf("Received chat message from client %d: %s\n", clientID, p)

		var msg ChatMessage
		if err := json.Unmarshal(p, &msg); err != nil {
			log.Println("Failed to parse message:", err)
			continue
		}

		messageChan <- msg

	}

	close(messageChan)
}

// processChatMessages processes chat messages from the channel
func processChatMessages(clientID int, messageChan chan ChatMessage) {
	for msg := range messageChan {
		if err := insertChatMessage(msg); err != nil {
			log.Printf("Failed to insert chat message from client %d: %v\n", clientID, err)
		} else {
			log.Printf("Inserted chat message from %s to %s: %s\n", msg.Sender, msg.Recipient, msg.Message)
		}
	}
	log.Printf("Message handler for client %d exiting\n", clientID)
}

// insertChatMessage inserts a chat message into the database
func insertChatMessage(msg ChatMessage) error {
	timestamp, err := time.Parse(time.RFC3339, msg.Timestamp)
	if err != nil {
		log.Printf("Invalid timestamp format: %v\n", err)
		return err
	}

	query := `
        INSERT INTO messages (project_id, part_id, message, sender, recipient, timestamp, user_type)
        VALUES (?, ?, ?, ?, ?, ?, ?)
    `
	_, err = database.DB.Exec(query, msg.ProjectID, msg.PartID, msg.Message, msg.Sender, msg.Recipient, timestamp, "engineer")
	return err
}
