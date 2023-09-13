package player

import (
	"context"
	"encoding/json"
	"github.com/a-h/templ"
	"github.com/gorilla/websocket"
	"gogame/htmx"
	"log"
)

type Player struct {
	conn        *websocket.Conn
	DisplayName string
	Send        chan templ.Component
}

func NewPlayer(conn *websocket.Conn, name string) Player {
	return Player{
		conn:        conn,
		DisplayName: name,
		Send:        make(chan templ.Component),
	}
}

type Message struct {
	Type    string          `json:"type"`
	Headers htmx.Headers    `json:"HEADERS"`
	Data    json.RawMessage `json:"data"`
}

type PlayerMsg struct {
	Player  *Player
	Message Message
}

func (p *Player) ListenForMessages(send chan PlayerMsg) {
	defer func() {
		p.conn.Close()
	}()
	for {
		_, msgBytes, err := p.conn.ReadMessage()
		if err != nil {
			// TODO: remove if else once how different errors are handled is decided
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("close error: %v\n", err)
			} else {
				log.Printf("error: %v\n", err)
			}
			break
		}
		var msgJson Message
		if err = json.Unmarshal(msgBytes, &msgJson); err != nil {
			log.Println("error: could not decode message:", string(msgBytes), "as a PlayerMsg")
		}
		msg := PlayerMsg{Player: p, Message: msgJson}
		send <- msg
	}
}

func (p *Player) WriteMessages() {
	defer func() {
		p.conn.Close()
	}()
	for {
		partial, ok := <-p.Send
		// channel closed
		if !ok {
			p.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		w, err := p.conn.NextWriter(websocket.TextMessage)
		if err != nil {
			log.Println(err)
			return
		}
		partial.Render(context.TODO(), w)

		if err := w.Close(); err != nil {
			log.Println(err)
			return
		}
	}
}
