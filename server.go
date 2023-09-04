package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	var count int = 0
	// Set routing rules
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// upgrade this connection to a WebSocket
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
		}

		log.Println("Client Connected")
		for {
			// read in a message
			messageType, p, err := ws.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			// print out that message for clarity
			log.Println("WebSocket Message Received: [", messageType,"]", string(p))

            count++;
			msg := []byte(fmt.Sprintf("<div id=\"count\">%d</div>", count))
			err = ws.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println(err)
				return
			}
		}
	})

	fmt.Println("Server started at 8080 port")
	//Use the default DefaultServeMux.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
