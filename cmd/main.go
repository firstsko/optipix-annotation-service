package main

import (
	"annotation-service/internal/annotation"
	"annotation-service/internal/chat"
	"annotation-service/pkg/config"
	"annotation-service/pkg/database"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.LoadConfig("configs/config.yaml")

	database.InitDatabase()

	http.HandleFunc("/ws/annotation", annotation.HandleConnections)
	http.HandleFunc("/ws/chat", chat.HandleConnections)

	serverConfig := config.AppConfig.Server
	address := fmt.Sprintf(":%d", serverConfig.Port)
	log.Printf("Server started on %s with log level %s", address, serverConfig.LogLevel)
	log.Fatal(http.ListenAndServe(address, nil))
}
