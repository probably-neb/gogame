package tictactoe

import (
    "net/http"
    "github.com/a-h/templ"
	"github.com/gorilla/websocket"
    "log"
    "encoding/json"
    "gogame/htmx"
    "context"
)

var upgrader = websocket.Upgrader {
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Move struct {
    cell string
    playerId int
}

type Player struct {
    symbol rune
    conn *websocket.Conn
}

type Game struct {
    host Player
    guest Player
    started bool
}

func listenForMoves(id int, conn *websocket.Conn, moves chan Move) {
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
        log.Println("request for", req.Headers.HXTriggerName,cell_id)
        moves <- Move{cell: cell_id, playerId: id}
    }
}

func sendMove(symbol rune, cell string, conn *websocket.Conn) {
    writer, err := conn.NextWriter(websocket.TextMessage)
    if err != nil {
        log.Println(err)
    }

    Box(cell, &symbol).Render(context.TODO(), writer)
    err = writer.Close()
    if err != nil {
        log.Println(err)
        return
    }
}

func (g *Game) Run() {
    moves := make(chan Move)
    go listenForMoves(0, g.host.conn, moves)
    go listenForMoves(1, g.guest.conn, moves)
    for {
        move := <-moves
        var symbol rune
        if move.playerId == 0 {
            symbol = g.host.symbol
        } else {
            symbol = g.guest.symbol
        }

        sendMove(symbol, move.cell, g.host.conn)
        sendMove(symbol, move.cell, g.guest.conn)
    }
}

func (g *Game) Start(host *websocket.Conn) {
    g.host.conn = host
    g.host.symbol = 'x'
    g.guest.conn = nil
    g.guest.symbol = 'o'
    g.started = true
}

func (g *Game) Join(guest *websocket.Conn) {
    // TODO: check if game is started & !full
    g.guest.conn = guest
}

type TTTRequest struct {
    Headers htmx.HXHeaders `json:"HEADERS"`
}

func AddHandlers() {
    http.Handle("/games/tictactoe", templ.Handler(Page()))
    var game = Game{started: false}
    http.HandleFunc("/games/tictactoe/ws", func(w http.ResponseWriter, r *http.Request) {
		// upgrade this connection to a WebSocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
		}
		log.Println("Client Connected")
        if !game.started {
            log.Println("starting game")
            game.Start(conn)
        } else {
            log.Println("joining game")
            game.Join(conn)
            log.Println("running game")
            game.Run()
        }
	})
}
