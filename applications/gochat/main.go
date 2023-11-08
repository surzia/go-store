package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Message struct {
	Username string   `json:"username"`
	Message  string   `json:"message"`
	Type     string   `json:"type"`
	ToUser   string   `json:"toUser"`
	Users    []string `json:"users"`
}

var clients = make(map[*websocket.Conn]string)
var broadcast = make(chan Message)
var history []Message
var onlineUsers []string
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var mutex sync.Mutex

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	var username string
	if name := r.URL.Query().Get("username"); name != "" {
		username = name
		clients[ws] = username // 将连接与用户名关联
		mutex.Lock()
		onlineUsers = append(onlineUsers, username)
		sendOnlineUsersList()
		mutex.Unlock()
	} else {
		log.Printf("Username not provided.")
		return
	}

	mutex.Lock()
	for _, msg := range history {
		err := ws.WriteJSON(msg)
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
	}
	mutex.Unlock()

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			mutex.Lock()
			delete(clients, ws)
			onlineUsers = removeUser(onlineUsers, username)
			sendOnlineUsersList()
			mutex.Unlock()
			break
		}

		if msg.Type == "message" {
			if msg.Username != "" && msg.Message != "" {
				if msg.ToUser != "" {
					sendMessageToUser(msg)
				} else {
					broadcast <- msg
				}
			}
		}
	}
}

func sendMessageToUser(msg Message) {
	for client, user := range clients {
		if user == msg.ToUser || user == msg.Username {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				mutex.Lock()
				delete(clients, client)
				onlineUsers = removeUser(onlineUsers, user)
				mutex.Unlock()
			}
		}
	}
}

func sendOnlineUsersList() {
	userListMessage := Message{
		Type:  "userlist",
		Users: onlineUsers,
	}

	for client := range clients {
		err := client.WriteJSON(userListMessage)
		if err != nil {
			log.Printf("error: %v", err)
			client.Close()
			mutex.Lock()
			delete(clients, client)
			mutex.Unlock()
		}
	}
}

func removeUser(users []string, username string) []string {
	for i, user := range users {
		if user == username {
			return append(users[:i], users[i+1:]...)
		}
	}
	return users
}

func handleMessages() {
	for {
		msg := <-broadcast

		mutex.Lock()
		history = append(history, msg)
		mutex.Unlock()

		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				mutex.Lock()
				delete(clients, client)
				mutex.Unlock()
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
