package tictactoe

import (
    "net/http"
    "github.com/a-h/templ"
	"github.com/gorilla/websocket"
    "log"
    "encoding/json"
    "gogame/htmx"
)

var upgrader = websocket.Upgrader {
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Player struct {
    symbol rune
    conn *websocket.Conn
}

var players [2]Player

type TTTRequest struct {
    Headers htmx.HXHeaders `json:"HEADERS"`
}

func AddHandlers() {
    http.Handle("/games/tictactoe", templ.Handler(Page()))
	http.HandleFunc("/games/tictactoe/ws", func(w http.ResponseWriter, r *http.Request) {
		// upgrade this connection to a WebSocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
		}
		log.Println("Client Connected")
        symbol := 'x'
        for {
            _, p, err := conn.ReadMessage()
            if err != nil {
                log.Println(err)
                return
            }
            var req TTTRequest
            if err = json.Unmarshal(p, &req); err != nil {
                log.Println(err)
                return
            }
            cell_id := req.Headers.HXTarget
            if err != nil {
                log.Println(err)
            }
            log.Println("request for", req.Headers.HXTriggerName,cell_id)
            writer, err := conn.NextWriter(websocket.TextMessage)
            if err != nil {
                log.Println(err)
            }

            Box(cell_id, symbol).Render(r.Context(), writer)
            err = writer.Close()
            if err != nil {
                log.Println(err)
                return
            }
            if symbol == 'x' {
                symbol = 'o'
            } else {
                symbol = 'x'
            }
        }
	})
}
