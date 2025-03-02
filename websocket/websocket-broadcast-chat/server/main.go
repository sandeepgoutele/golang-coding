package main

import (
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]string)
var broadcast = make(chan []byte)
var mu sync.Mutex

func isClientInfo(msg []byte) bool {
	return !strings.Contains(string(msg), ":")
}

func getClientIDFromMsg(msg []byte) string {
	// Assuming the message contains the client ID in a specific format, e.g., "clientID:message"
	parts := strings.SplitN(string(msg), ":", 2)
	if len(parts) > 1 {
		return parts[0]
	}
	return ""
}

func getTextFromMsg(msg []byte) string {
	// Assuming the message contains the client ID in a specific format, e.g., "clientID:message"
	parts := strings.SplitN(string(msg), ":", 2)
	if len(parts) > 1 {
		return parts[1]
	}
	return ""
}

func printClientInfo(msg []byte) {
	clientID := getClientIDFromMsg(msg)
	log.Printf("Message from client (%s): %s", clientID, getTextFromMsg(msg))
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Error upgrading connection: ", err)
	}
	defer conn.Close()

	remoteAdd := strings.Split(conn.RemoteAddr().String(), ":")
	clientId := remoteAdd[len(remoteAdd)-1]
	log.Println("New client connected with cliend id: ", clientId)

	for {
		// Read message from the client
		_, msg, err := conn.ReadMessage()
		if err != nil {
			mu.Lock()
			delete(clients, conn)
			mu.Unlock()
			log.Println("Client disconnected.")
			break
		}

		// Print the message to the console
		// log.Printf("Message from client (%s): %s", clientId, msg)

		// Write message back to the client
		// err = conn.WriteMessage(websocket.TextMessage, msg)
		// if err != nil {
		// log.Println("Error writing message: ", err)
		// break
		// }
		if isClientInfo(msg) {
			mu.Lock()
			clients[conn] = string(msg)
			mu.Unlock()
			log.Printf("Client ID %s communicated.", string(msg))
		} else {
			printClientInfo(msg)
			broadcast <- msg
			if getTextFromMsg(msg) == "exit" {
				mu.Lock()
				clientId := getClientIDFromMsg(msg)
				for client := range clients {
					if clients[client] == clientId {
						delete(clients, client)
						client.Close()
						log.Printf("Client %s disconnected.", clientId)
						break
					}
				}
				mu.Unlock()
				break
			}
		}
	}
}

func handleMessages() {
	for {
		// Get the message from the broadcast channel
		msg := <-broadcast
		clientId := getClientIDFromMsg(msg)
		// Send message to all clients
		mu.Lock()
		for client := range clients {
			if clients[client] == clientId {
				continue
			}
			err := client.WriteMessage(websocket.TextMessage, []byte(getTextFromMsg(msg)+" (from: "+clientId+")"))
			if err != nil {
				log.Println("Error writing message: ", err)
				client.Close()
				delete(clients, client)
			}
		}
		mu.Unlock()
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()
	log.Println("WebSocket Server started on ws://localhost:8081/ws")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}
