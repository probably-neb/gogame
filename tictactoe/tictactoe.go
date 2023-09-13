package tictactoe

import (
	"encoding/json"
	"gogame/htmx"
	"gogame/player"
	"log"
)

type Player = player.Player

type Symbol rune

const (
	X Symbol = 'X'
	O        = 'O'
)

type Move struct {
	cell   string
	player *Player
}

type TTTRequest struct {
	Headers htmx.Headers `json:"HEADERS"`
}

func listenForMoves(msgs chan player.PlayerMsg, moves chan Move) {
	for {
		// TODO: check if channel closed
		msg := <-msgs

		// TODO: assert/check msg.Type
		var req TTTRequest
		if err := json.Unmarshal(msg.Message.Data, &req); err != nil {
			log.Println(err)
			return
		}
		cell_id := msg.Message.Headers.HXTarget
		log.Println(msg.Player, "played", cell_id)
		moves <- Move{cell: cell_id, player: msg.Player}
	}
}

func Start(host *Player, guest *Player, msgs chan player.PlayerMsg, exit chan struct{}) {
    defer func() {exit <- struct{}{}}()
	moves := make(chan Move)
	symbols := map[*Player]Symbol{host: X, guest: O}
	host.Send <- Game()
	guest.Send <- Game()
	go listenForMoves(msgs, moves)
	turn := X
	for {
		move := <-moves
		symbol := symbols[move.player]
		if symbol != turn {
			log.Println("wrong turn it's", turn, "turn")
			// TODO: consider send not turn warning?
			continue
		} else if turn == symbols[host] {
			log.Println("guest's turn")
			turn = symbols[guest]
		} else {
			log.Println("host's turn")
			turn = symbols[host]
		}
        symbolRune := rune(symbol)
		newBox := Box(move.cell, &symbolRune)
		host.Send <- newBox
		guest.Send <- newBox
	}
}
