package main

import (
	"bufio"
	"log"
	"os"
	"os/signal"

	"fmt"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func generateClientID() string {
	return uuid.New().String()
}
func main() {
	clientID := generateClientID()
	fmt.Printf("Generated Client ID: %s\n", clientID)

	serverUrl := "ws://localhost:8081/ws"
	conn, _, err := websocket.DefaultDialer.Dial(serverUrl, nil)
	if err != nil {
		log.Fatal("Failed to connect to WebSocket server with error: ", err)
	}
	defer conn.Close()

	err = conn.WriteMessage(websocket.TextMessage, []byte(clientID))
	if err != nil {
		log.Fatal("Failed to send client ID with error: ", err)
	}

	log.Printf("Connected to WebSocket server. Type your message and press Enter to send.")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message: ", err)
				return
			}
			log.Printf("Received message: %s", message)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		textMsg := scanner.Text()
		msg := clientID + ":" + textMsg
		if textMsg == "" {
			continue
		}
		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Println("Error writing message: ", err)
			return
		}
		if textMsg == "exit" {
			log.Print("Exiting WebSocket client...")
			break
		}
	}

	closeMsg := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Client exiting")
	err = conn.WriteMessage(websocket.CloseMessage, closeMsg)
	if err != nil {
		log.Println("Error writing close message: ", err)
		return
	}
}
