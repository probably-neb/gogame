package tictactoe

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
    "gogame/htmx"
	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Symbol rune

type Move struct {
    cell string
    symbol Symbol
}

type TTTPlayer struct {
    symbol Symbol
    conn *websocket.Conn
}

type TicTacToe struct {
    host Player
    guest Player
    hostSym Symbol
    guestSym Symbol
}

func (g *Game) Start(host *websocket.Conn) {
    g.host.conn = host
    g.host.symbol = 'x'
    g.guest.conn = nil
    g.guest.symbol = 'o'
    g.started = true
}

func (g *Game) Run() {
    moves := make(chan Move)
    go listenForMoves(g.host.symbol, g.host.conn, moves)
    go listenForMoves(g.guest.symbol, g.guest.conn, moves)
    turn := g.host.symbol
    for {
        move := <-moves
        if move.symbol != turn {
            log.Println("wrong turn it's", turn, "turn")
            // TODO: send error
            continue;
        } else if turn == g.host.symbol {
            log.Println("guest's turn")
            turn = g.guest.symbol
        } else {
            log.Println("host's turn")
            turn = g.host.symbol
        }

        sendMove(move, g.host.conn)
        sendMove(move, g.guest.conn)
    }
}


func listenForMoves(symbol Symbol, conn *websocket.Conn, moves chan Move) {
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
        log.Println(string(symbol),"played",cell_id)
        moves <- Move{cell: cell_id, symbol: symbol}
    }
}

func sendMove(move Move, conn *websocket.Conn) {
    writer, err := conn.NextWriter(websocket.TextMessage)
    if err != nil {
        log.Println(err)
    }
    symbol := rune(move.symbol)
    Box(move.cell, &symbol).Render(context.TODO(), writer)
    err = writer.Close()
    if err != nil {
        log.Println(err)
        return
    }
}


type TTTRequest struct {
    Headers htmx.Headers `json:"HEADERS"`
}

func AddHandlers(r *mux.Router) {
    http.Handle("/games/tictactoe", templ.Handler(Page()))
    var game = Game{started: false}
    http.HandleFunc("/games/tictactoe/ws", func(w http.ResponseWriter, r *http.Request) {
		// upgrade this connection to a WebSocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
		}
		log.Println("Client Connected", conn.RemoteAddr(), conn.LocalAddr())
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
