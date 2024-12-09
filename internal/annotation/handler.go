package annotation

import (
	"annotation-service/pkg/database"
	"annotation-service/pkg/websocket"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Annotation struct {
	ID          int     `json:"id"`
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
	Width       float64 `json:"width"`
	Height      float64 `json:"height"`
	ImageWidth  int     `json:"imageWidth"`
	ImageHeight int     `json:"imageHeight"`
}

type AnnotationMessage struct {
	ProjectID   json.Number  `json:"project_id"`
	PartID      json.Number  `json:"part_id"`
	Annotations []Annotation `json:"annotations"`
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	clientID := websocket.GenerateClientID()
	log.Printf("New annotation client connected with ID: %d\n", clientID)

	upgrader := websocket.GetUpgrader()
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	messageChan := make(chan []byte)

	//Go routine
	go handleMessages(clientID, messageChan)

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		log.Printf("Received annotation message from client %d: %s\n", clientID, p)

		messageChan <- p
	}
	close(messageChan)
}

func handleMessages(clientID int, messageChan chan []byte) {
	for p := range messageChan {
		var msg AnnotationMessage
		if err := json.Unmarshal(p, &msg); err != nil {
			log.Printf("Failed to parse message from client %d: %v\n", clientID, err)
			continue
		}

		projectID, err := strconv.Atoi(msg.ProjectID.String())
		if err != nil {
			log.Printf("Invalid project_id from client %d: %v\n", clientID, err)
			continue
		}

		partID, err := strconv.Atoi(msg.PartID.String())
		if err != nil {
			log.Printf("Invalid part_id from client %d: %v\n", clientID, err)
			continue
		}

		for _, annotation := range msg.Annotations {
			err := insertAnnotation(projectID, partID, annotation)
			if err != nil {
				log.Printf("Failed to insert annotation from client %d: %v\n", clientID, err)
			} else {
				log.Printf("Inserted annotation from client %d: %+v\n", clientID, annotation)
			}
		}
	}
	log.Printf("Message handler for client %d exiting\n", clientID)
}

func insertAnnotation(projectID, partID int, annotation Annotation) error {
	query := `
        INSERT INTO annotations (project_id, part_id, x, y, width, height, image_width, image_height, timestamp)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, NOW())
    `
	_, err := database.DB.Exec(query, projectID, partID, annotation.X, annotation.Y, annotation.Width, annotation.Height, annotation.ImageWidth, annotation.ImageHeight)
	return err
}
