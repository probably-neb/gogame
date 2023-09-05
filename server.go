package main

import (
	"encoding/json"
	"fmt"
	"github.com/a-h/templ"
	"github.com/gorilla/websocket"
	"gogame/games/tictactoe"
	"gogame/partials"
    "gogame/htmx"
	"log"
	"net/http"
)

type HXWSCountMessage struct {
	Method  string    `json:"method"`
	Headers htmx.HXHeaders `json:"HEADERS"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var count int = 0
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)

func sendCount(conn *websocket.Conn) {
	response := []byte(fmt.Sprintf("<div id=\"count\">%d</div>", count))
	err := conn.WriteMessage(websocket.TextMessage, response)
	if err != nil {
		log.Println(err)
		delete(clients, conn)
		return
	}
}

func broadcastMessages() {
	for {
		// grab the next message from the broadcast channel
		cmd := <-broadcast
		log.Println("propogating:", cmd)
		switch cmd {
		case "increment":
			count++
		case "decrement":
			count--
		default:
			log.Println("Unknown method:", cmd)
		}

		// send it out to every client that is currently connected
		for conn := range clients {
			sendCount(conn)
		}
	}
}

func main() {
	// Set routing rules
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	go broadcastMessages()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// upgrade this connection to a WebSocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
		}
		clients[conn] = true
		sendCount(conn)

		log.Println("Client Connected")
		for {
			_, p, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				delete(clients, conn)
				return
			}
			var cmd HXWSCountMessage
			if err := json.Unmarshal(p, &cmd); err != nil {
				log.Println(err)
			}
			// print out that message for clarity
			log.Println("WebSocket Message Received:", cmd.Method)

			broadcast <- cmd.Method
		}
	})
	http.Handle("/board", templ.Handler(partials.Board(20, 20)))
	tictactoe.AddHandlers()

	fmt.Println("Server started at 8080 port")
	//Use the default DefaultServeMux.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
