package player

import (
    "github.com/gorilla/websocket"
    "log"
)


type Player struct {
	conn        *websocket.Conn
	DisplayName string
    Send        chan []byte
}

type PlayerMsg struct {
    Player *Player
    Message []byte
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
				log.Printf("close error: %v", err)
			} else {
                log.Printf("error: %v", err)
            }
			break
		}
        msg := PlayerMsg{Player: p, Message: msgBytes}
		send <- msg
	}
}

func (p *Player) WriteMessages() {
    defer func() {
        p.conn.Close()
    }()
    for {
		message, ok := <-p.Send;
            // channel closed
        if !ok {
            p.conn.WriteMessage(websocket.CloseMessage, []byte{})
            return
        }

        w, err := p.conn.NextWriter(websocket.TextMessage)
        if err != nil {
            return
        }
        w.Write(message)

        if err := w.Close(); err != nil {
            return
        }
	}
}
